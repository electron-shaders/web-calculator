//TODO: 检测负数
//TODO: 支持小数

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
	Error        error  `json:"error"`
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
	if regexp.MustCompile(`[a-zA-Z]`).MatchString(orig) {
		return nil, errors.New("the expression cannot contain letter(s)")
	} else if regexp.MustCompile(`[^\+\-\*/\(\)0-9]`).MatchString(orig) {
		return nil, errors.New("the expression contains invalid character(s)")
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
	temp := strings.Replace(string(tmp), " ", "", -1)
	temp = strings.Replace(temp, "\n", "", -1)
	res.CorrectedExp = temp
	fmt.Println("修正结果:", temp)
	indexs, err := findIndOfOps(temp)
	if err != nil {
		res.Error = err
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
			x, _ := strconv.Atoi(parser.Pop())
			y, _ := strconv.Atoi(parser.Pop())
			switch parsedExp[i] {
			case "+":
				parser.Push(strconv.Itoa(y + x))
			case "-":
				parser.Push(strconv.Itoa(y - x))
			case "*":
				parser.Push(strconv.Itoa(y * x))
			case "/":
				if x == 0 {
					return 0, errors.New("the divider cannot be zero")
				} else {
					parser.Push(strconv.Itoa(y / x))
				}
			}
		}
	}
	ans, err := strconv.Atoi(parser.Pop())
	if err != nil {
		panic(err)
	}
	return ans, nil
}

func main() {
	defer panicHandler()
	http.HandleFunc("/", process)
	http.ListenAndServe(":3001", nil)
}

func process(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var req ReqData
	if err := decoder.Decode(&req); err != nil {
		panic(err)
	}

	if err := preParse(req.Tmp); err != nil {
		fmt.Println("错误:", err)
		res.Error = err
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusBadRequest)
		if jsonErr := json.NewEncoder(w).Encode(res); jsonErr != nil {
			panic(jsonErr)
		}
	}

	if ans, err := calc(); err != nil {
		fmt.Println("错误:", err)
		res.Error = err
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusBadRequest)
		if jsonErr := json.NewEncoder(w).Encode(res); jsonErr != nil {
			panic(jsonErr)
		}
	} else {
		fmt.Println("计算结果:", ans)
		res.Answer = ans
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		if jsonErr := json.NewEncoder(w).Encode(res); jsonErr != nil {
			panic(jsonErr)
		}
	}
}

func panicHandler() {
	err := recover()
	if err != nil {
		fmt.Println("Panic:", err)
	}
}
