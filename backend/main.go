//TODO: 检测负数
//TODO: 支持小数
//FIXME: 操作数过长

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/electron-shaders/web-calculator/backend/stack"
)

type ResData struct {
	Answer       int    `json:"answer"`
	CorrectedExp string `json:"corrected-exp"`
	Error        string `json:"error-msg"`
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
	if regexp.MustCompile(`[^\+\-\*/\(\)0-9]`).MatchString(orig) {
		return nil, errors.New("表达式包含非法字符")
	}
	return regexp.MustCompile(`(\+|\-|\*|/|\(|\))`).FindAllStringIndex(orig, -1), nil
}

func oplv(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "(", ")":
		return 3
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
	res.CorrectedExp = temp
	fmt.Println("修正结果:", temp)
	indexs, err := findIndOfOps(temp)
	if err != nil {
		res.Error = err.Error()
		return err
	}
	if len(indexs) > 0 {
		for i := 0; i < len(temp); i++ {
			if i != indexs[tot][0] {
				parsedExp += string(temp[i])
			} else if string(temp[i]) == "(" {
				parsedExp += string(temp[i]) + " "
				if tot < len(indexs)-1 {
					tot++
				}
			} else if string(temp[i]) == ")" {
				parsedExp += " " + string(temp[i])
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
	return nil
}

func calc() (int, error) {
	origExp = strings.Split(parsedExp, " ")
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
			} else if oplv(parser.Top()) <= oplv(origExp[i]) {
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
			x, err := strconv.Atoi(parser.Pop())
			if err != nil {
				return 0, errors.New("操作数过大")
			}
			y, err := strconv.Atoi(parser.Pop())
			if err != nil {
				return 0, errors.New("操作数过大")
			}
			switch parsedExp[i] {
			case "+":
				parser.Push(strconv.Itoa(y + x))
			case "-":
				parser.Push(strconv.Itoa(y - x))
			case "*":
				parser.Push(strconv.Itoa(y * x))
			case "/":
				if x == 0 {
					return 0, errors.New("除数不可为 0")
				} else {
					parser.Push(strconv.Itoa(y / x))
				}
			}
		}
	}
	if ans, err := strconv.Atoi(parser.Pop()); err != nil {
		panic(err)
	} else {
		return ans, nil
	}
}

func main() {
	http.HandleFunc("/", process)
	http.ListenAndServe(":3001", nil)
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
		fmt.Println("错误:", err)
		res.Error = err.Error()
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusOK)
		if jsonErr := json.NewEncoder(w).Encode(res); jsonErr != nil {
			panic(jsonErr)
		}
	} else if ans, err := calc(); err != nil {
		fmt.Println("错误:", err)
		res.Error = err.Error()
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.WriteHeader(http.StatusOK)
		if jsonErr := json.NewEncoder(w).Encode(res); jsonErr != nil {
			panic(jsonErr)
		}
	} else {
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
}
