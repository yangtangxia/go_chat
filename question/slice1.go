package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	x := 2
	y := 4
	table := make([][]int, x)
	for i := range table {
		table[i] = make([]int, y)
	}
	fmt.Println(table)

	h, w := 2, 4
	raw := make([]int, h*w)
	for i := range raw {
		raw[i] = i
	}
	fmt.Println(raw)

	for i := range table {
		table[i] = raw[i*w : i*w+w]
	}

	fmt.Println(table)

	// 正确判断某个键是否存在
	s := map[string]string{"one": "1", "two": ""}
	if _, ok := s["www"]; !ok {
		fmt.Println("key two is no entry")
	}

	// 正确修改字符串
	str := "text"
	// str[0] = "T" // 错误赋值修改
	// fmt.Println(str)

	// 正确修改字符串
	strByter := []byte(str)
	strByter[0] = 'T' // 这里只能是单引号，不能双引号，否在报类型错误
	str = string(strByter)
	fmt.Println(str)

	// 中文替换：因为中文占3-4个字符
	strRunes := []rune(str)
	strRunes[0] = '我'
	str = string(strRunes)
	fmt.Println(str)

	str1 := "ABC"
	fmt.Println(utf8.ValidString(str1)) // 判断字符串是不是utf8文本
	str2 := "A\\xfeC"
	fmt.Println(utf8.ValidString(str2))

	fmt.Println(str2)

	// 字符串长度
	char := "♥"
	fmt.Println(len(char))

	fmt.Println(utf8.RuneCountInString(char))

	char1 := "é"
	fmt.Println(len(char1))

	fmt.Println(utf8.RuneCountInString(char1))

	xx := []int{
		1,
		2, // 结尾逗号必须要
	}

	yy := []int{1, 2}

	zz := []int{1, 2}

	fmt.Println(xx, yy, zz)
}
