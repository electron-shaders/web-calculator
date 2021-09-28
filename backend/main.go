package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/electron-shaders/web-calculator/backend/stack"
)

type ResData struct {
	Answer       float64 `json:"answer"`
	CorrectedExp string  `json:"corrected-exp"`
	Error        string  `json:"error-msg"`
}

func (res *ResData) clear() {
	res.Answer = 0
	res.CorrectedExp = ""
	res.Error = ""
}

type ReqData struct {
	Tmp string `json:"orig-exp"`
}

var (
	res       ResData
	origExp   []string
	parsedExp string
	parser    stack.StringStack
)

func findIndOfOps(orig string) ([][]int, error) {
	if regexp.MustCompile(`[^\+\-\*/\(\)0-9\.\^#]`).MatchString(orig) {
		return nil, errors.New("表达式包含非法字符")
	}
	return regexp.MustCompile(`(\+|\-|\*|/|\(|\)|\^|#)`).FindAllStringIndex(orig, -1), nil
}

func oplv(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^", "#":
		return 3
	case "(", ")":
		return 4
	default:
		return -1
	}
}

func preParse(tmp string) error {
	var tot int
	if len(tmp) == 0 {
		return errors.New("表达式不可为空")
	}
	temp := strings.Replace(tmp, " ", "", -1)
	temp = strings.Replace(temp, "\n", "", -1)
	temp = strings.Replace(temp, "。", ".", -1)
	temp = strings.Replace(temp, "（", "(", -1)
	temp = strings.Replace(temp, "）", ")", -1)
	temp = strings.Replace(temp, "、", "/", -1)
	res.CorrectedExp = temp
	fmt.Println("修正结果:", temp)
	temp = strings.Replace(temp, "√", "#", -1)
	temp = strings.Replace(temp, "%", "/100", -1)
	indexs, err := findIndOfOps(temp)
	if err != nil {
		res.Error = err.Error()
		return err
	}
	bracket := 0
	if len(indexs) > 0 {
		for i := 0; i < len(temp); i++ {
			if i != indexs[tot][0] {
				parsedExp += string(temp[i])
			} else if string(temp[i]) == "#" {
				parsedExp += "# "
				if tot < len(indexs)-1 {
					tot++
				}
			} else if string(temp[i]) == "(" {
				bracket++
				parsedExp += "( "
				if tot < len(indexs)-1 {
					tot++
				}
			} else if string(temp[i]) == ")" {
				bracket--
				parsedExp += " )"
				if tot < len(indexs)-1 {
					tot++
				}
			} else {
				parsedExp += " " + string(temp[i]) + " "
				if tot < len(indexs)-1 {
					tot++
				}
			}
		}
	} else {
		for i := 0; i < len(temp); i++ {
			parsedExp += string(temp[i])
		}
	}
	if bracket > 0 {
		return errors.New("括号匹配不完整")
	}
	return nil
}

func calc() (float64, error) {
	origExp = strings.Split(parsedExp, " ")
	for i := 0; i < len(origExp); i++ {
		if origExp[i] == "-" && origExp[i-1] == "" {
			origExp[i] = "-" + origExp[i+1]
			origExp = append(origExp[:i-1], append(origExp[i:i+1], origExp[i+2:]...)...)
		}
	}
	var parsedExp []string
	for i := 0; i < len(origExp); i++ {
		if oplv(origExp[i]) == -1 {
			parsedExp = append(parsedExp, string(origExp[i]))
		} else {
			if origExp[i] == ")" {
				for parser.Top() != "(" {
					parsedExp = append(parsedExp, parser.Pop())
				}
				parser.Pop()
			} else if oplv(parser.Top()) < oplv(origExp[i]) {
				parser.Push(origExp[i])
			} else {
				for oplv(parser.Top()) >= oplv(origExp[i]) && parser.Top() != "(" {
					parsedExp = append(parsedExp, parser.Pop())
				}
				parser.Push(origExp[i])
			}
		}
	}
	for !parser.IsEmpty() {
		parsedExp = append(parsedExp, parser.Pop())
	}
	origExp = nil
	for i := 0; i < len(parsedExp); i++ {
		if oplv(parsedExp[i]) == -1 {
			parser.Push(parsedExp[i])
		} else {
			x, _ := strconv.ParseFloat(parser.Pop(), 64)
			y, _ := strconv.ParseFloat(parser.Pop(), 64)
			switch parsedExp[i] {
			case "+":
				parser.Push(fmt.Sprintf("%.9f", y+x))
			case "-":
				parser.Push(fmt.Sprintf("%.9f", y-x))
			case "*":
				parser.Push(fmt.Sprintf("%.9f", y*x))
			case "/":
				if x == 0 {
					return 0, errors.New("除数不可为 0")
				} else {
					parser.Push(fmt.Sprintf("%.9f", y/x))
				}
			case "^":
				parser.Push(fmt.Sprintf("%.9f", math.Pow(y, x)))
			case "#":
				parser.Push(fmt.Sprintf("%.9f", math.Sqrt(x)))
			}
		}
	}
	if ans, err := strconv.ParseFloat(parser.Pop(), 64); err == nil {
		return ans, nil
	} else if regexp.MustCompile(`invalid syntax`).MatchString(err.Error()) {
		return 0, nil
	} else {
		fmt.Println(err)
		return 0, errors.New("操作数过大")
	}
}

func main() {
	http.HandleFunc("/", process)
	http.ListenAndServe(":3001", nil)
}

func sendErr(w http.ResponseWriter, err error) {
	fmt.Println("错误:", err)
	res.Error = err.Error()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusOK)
	if jsonErr := json.NewEncoder(w).Encode(res); jsonErr != nil {
		panic(jsonErr)
	}
}

func sendAns(w http.ResponseWriter, ans float64) {
	fmt.Println("计算结果:", ans)
	res.Answer = ans
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusOK)
	if jsonErr := json.NewEncoder(w).Encode(res); jsonErr != nil {
		panic(jsonErr)
	}
}

func process(w http.ResponseWriter, request *http.Request) {
	parser.Clear()
	res.clear()
	origExp = nil
	parsedExp = ""
	decoder := json.NewDecoder(request.Body)
	var req ReqData
	if err := decoder.Decode(&req); err != nil {
		panic(err)
	}
	if err := preParse(req.Tmp); err != nil {
		sendErr(w, err)
	} else if ans, err := calc(); err != nil {
		sendErr(w, err)
	} else {
		sendAns(w, ans)
	}
}
