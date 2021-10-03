package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/electron-shaders/web-calculator/backend/color"
	"github.com/electron-shaders/web-calculator/backend/stack"
	"github.com/gin-gonic/gin"
)

type ReqData struct {
	Tmp string `json:"orig-exp"`
}

var (
	correctedExp string
	origExp      []string
	parsedExp    string
	parser       stack.StringStack
)

func isInvalidSyntax(err error) bool {
	return regexp.MustCompile(`invalid syntax`).MatchString(err.Error())
}

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
	correctedExp = temp
	timeStr := time.Now().Format("2006/01/02 - 03:04:05")
	fmt.Printf("[INFO] %s 修正结果: %s\n", timeStr, temp)
	temp = strings.Replace(temp, "√", "#", -1)
	temp = strings.Replace(temp, "%", "/100", -1)
	indexs, err := findIndOfOps(temp)
	if err != nil {
		return err
	}
	bracket := 0
	if len(indexs) > 0 {
		for i := 0; i < len(temp); i++ {
			if i != indexs[tot][0] {
				parsedExp += string(temp[i])
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
			switch parsedExp[i] {
			case "+":
				y, _ := strconv.ParseFloat(parser.Pop(), 64)
				parser.Push(fmt.Sprintf("%.9f", y+x))
			case "-":
				y, _ := strconv.ParseFloat(parser.Pop(), 64)
				parser.Push(fmt.Sprintf("%.9f", y-x))
			case "*":
				y, _ := strconv.ParseFloat(parser.Pop(), 64)
				parser.Push(fmt.Sprintf("%.9f", y*x))
			case "/":
				y, _ := strconv.ParseFloat(parser.Pop(), 64)
				if x == 0 {
					return 0, errors.New("除数不可为 0")
				} else {
					parser.Push(fmt.Sprintf("%.9f", y/x))
				}
			case "^":
				y, _ := strconv.ParseFloat(parser.Pop(), 64)
				ans := math.Pow(y, x)
				if math.IsNaN(ans) {
					return 0, errors.New("算术错误")
				}
				parser.Push(fmt.Sprintf("%.9f", ans))
			case "#":
				y, err := strconv.ParseFloat(parser.Pop(), 64)
				if err != nil && isInvalidSyntax(err) {
					y = 1
				}
				if x < 0 {
					return 0, errors.New("平方根的被开方数不可为负数")
				}
				parser.Push(fmt.Sprintf("%.9f", y*math.Sqrt(x)))
			}
		}
	}
	if ans, err := strconv.ParseFloat(parser.Pop(), 64); err == nil {
		return ans, nil
	} else if isInvalidSyntax(err) {
		return 0, nil
	} else {
		fmt.Println(err)
		return 0, errors.New("操作数过大")
	}
}

func main() {
	router := gin.Default()
	router.POST("/process", func(c *gin.Context) {
		parser.Clear()
		correctedExp = ""
		origExp = nil
		parsedExp = ""
		var req ReqData
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusOK, gin.H{"error-msg": err.Error()})
			fmt.Println(err)
			return
		}
		if err := preParse(req.Tmp); err != nil {
			timeStr := time.Now().Format("2006/01/02 - 03:04:05")
			fmt.Println("[ERR]", timeStr, "错误:", color.Color(color.FgWhite, color.BgRed, err.Error()))
			c.JSON(http.StatusOK, gin.H{
				"answer":        nil,
				"corrected-exp": nil,
				"error-msg":     err.Error(),
			})
		} else if ans, err := calc(); err != nil {
			timeStr := time.Now().Format("2006/01/02 - 03:04:05")
			fmt.Println("[ERR]", timeStr, "错误:", color.Color(color.FgWhite, color.BgRed, err.Error()))
			c.JSON(http.StatusOK, gin.H{
				"answer":        nil,
				"corrected-exp": nil,
				"error-msg":     err.Error(),
			})
		} else {
			finalAns := ""
			if math.IsInf(ans, 1) {
				finalAns = "∞"
			} else if math.IsInf(ans, -1) {
				finalAns = "-∞"
			}
			timeStr := time.Now().Format("2006/01/02 - 03:04:05")
			if finalAns != "" {
				fmt.Println("[INFO]", timeStr, "计算结果:", finalAns)
				c.JSON(http.StatusOK, gin.H{
					"answer":        finalAns,
					"corrected-exp": correctedExp,
					"error-msg":     nil,
				})
			} else {
				fmt.Println("[INFO]", timeStr, "计算结果:", ans)
				c.JSON(http.StatusOK, gin.H{
					"answer":        ans,
					"corrected-exp": correctedExp,
					"error-msg":     nil,
				})
			}
		}
	})
	router.Run(":3001")
}
