//TODO: 检测负数
//TODO: 支持小数

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/electron-shaders/web-calculator/backend/stack"
)

var (
	origExp             []string
	parsedExp           string
	parser              stack.StringStack
	isWarnedEmptyNumber bool
)

func findIndOfOps(orig string) ([][]int, error) {
	if regexp.MustCompile(`[a-zA-Z]`).MatchString(orig) {
		return nil, errors.New("表达式中不可包含字母")
	} else if regexp.MustCompile(`[^\+\-\*/\(\)0-9]`).MatchString(orig) {
		return nil, errors.New("表达式中包含非法字符")
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

func init() {
	var tot int
	//tmp := "+-*"
	print("请输入一个表达式(不支持小数，仅支持四则运算): ")
	reader := bufio.NewReader(os.Stdin)
	tmp, _, err := reader.ReadLine()
oserror:
	for err != nil {
		println()
		fmt.Println("错误:", err)
		print("请输入一个表达式(不支持小数，仅支持四则运算): ")
		tmp, _, err = reader.ReadLine()
	}
	temp := strings.Replace(string(tmp), " ", "", -1)
	temp = strings.Replace(temp, "\n", "", -1)
	fmt.Println("修正结果:", temp)
	indexs, err := findIndOfOps(temp)
	for err != nil {
		println()
		fmt.Println("错误:", err)
		print("请输入一个表达式(不支持小数，仅支持四则运算): ")
		tmp, _, err = reader.ReadLine()
		if err != nil {
			goto oserror
		}
		temp = strings.Replace(string(tmp), " ", "", -1)
		temp = strings.Replace(temp, "\n", "", -1)
		indexs, err = findIndOfOps(temp)
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
}

func main() {
	defer handlePanic()
	origExp = strings.Split(parsedExp, " ")
	var parsedExp []string
	for i := 0; i < len(origExp); i++ {
		if oplv(origExp[i]) == -1 {
			if origExp[i] == "" && !isWarnedEmptyNumber {
				println()
				fmt.Println("警告: 表达式不完整(缺省值：0)")
				isWarnedEmptyNumber = true
			} else {
				parsedExp = append(parsedExp, string(origExp[i]))
			}
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
					panic("除数为0")
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
	fmt.Println("计算结果:", ans)
}

func handlePanic() {
	err := recover()
	if err != nil {
		println()
		fmt.Println("致命错误:", err)
		fmt.Println("程序已中止...")
	}
}
