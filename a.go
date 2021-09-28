package main

import "fmt"

func main() {
	//len()无法正常直接计算包含汉字的string的长度
	str := "√(0)+√(0)"
	fmt.Println("str的长度：", len(str))

	//for-range获取的下标(index)无法用来正常打印string中的汉字
	for ind := range str {
		fmt.Printf("%c", str[ind])
	}
	println()

	//但可以用for-range获取的值(val)来正常打印
	for _, ch := range str {
		fmt.Printf("%c", ch)
	}
	println()

	//[]rune(str)可以正常计算包含汉字的string的长度，也可以正常使用str[index]
	str1 := []rune(str)
	fmt.Println("str1的长度为：", len(str1))
	for ind, ch := range str1 {
		fmt.Printf("str1[%d]=%c, ch=%c\n", ind, str1[ind], ch)
	}
}
