package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	str := " Beginning of the string " +
		"second part of the string "
	//print(str)
	//s := "hel" + "lo"
	//s += "world"
	//fmt.Println(s)
	//str1 := "asSASA DDD DS 你好 こん" // 中文算3个字符
	//fmt.Printf("str1 len is %d\n", len(str1))
	//fmt.Printf("characts is len %d\n", utf8.RuneCountInString(str1))
	//
	//res := strings.HasPrefix(str, "B")  // 判断字符串str是否以B开头 bool
	//fmt.Printf("%t\n", res)
	//row := strings.HasSuffix(str, "string") // 判断字符串str是否以string结尾 bool
	//fmt.Printf("%t\n", row)
	//
	//fmt.Printf("字符串是否包含%t\n", strings.Contains(str, "secosnd"))
	//
	//fmt.Printf("字符在字符串中最开始的索引%t\n", strings.Index(str, "u")) // int  -1未匹配到
	//fmt.Printf("字符在字符串中最后出现位置%t\n", strings.LastIndex(str, "s"))
	//fmt.Printf("字符在字符串中非ascll位置%t\n", strings.IndexRune(str1, rune('好')))
	//
	//fmt.Printf("字符串替换%t\n", strings.Replace(str, "string", "string", -1))
	//
	//fmt.Printf("字符串中出现次数%t\n", strings.Count(str, "string"))
	//
	//fmt.Printf("重复count次字符串%t\n", strings.Repeat(str, 2))
	//
	//fmt.Printf("修改字符串转化小写%t\n", strings.ToLower(str))
	//fmt.Printf("修改字符串转化大写%t\n", strings.ToUpper(str))
	//
	//fmt.Printf("去掉字符串开头结尾空格%t\n", strings.TrimSpace(str))
	//fmt.Printf("去掉字符串开头节气指定字符%t\n", strings.Trim(str, " "))
	//fmt.Printf("字符串分隔,默认使用了空白符号%t\n", strings.Fields(str))
	//fmt.Printf("字符串分隔%t\n", strings.Split(str, ","))

	s1 := strings.Fields(str)
	fmt.Printf("分隔字符串：%v\n", s1)
	for _, val := range s1{
		fmt.Printf("%s - ", val)
	}

	fmt.Println()
	str1 := "good |The Abc |go|25"
	s2 := strings.Split(str1, "|")
	fmt.Printf("分隔字符串%v\n", s2)
	for _, val := range s2 {
		fmt.Printf("%s - ", val)
	}

	fmt.Println()
	s3 := strings.Join(s2, ",")
	fmt.Printf("s2 连接 ：%s\n", s3)

	var orig string = "666"
	var an int
	var newS string
	fmt.Printf("the size of ints is : %d\n", strconv.IntSize)

	an, err := strconv.Atoi(orig)
	fmt.Printf("error: %s\n", err)
	fmt.Printf("the integer is %d\n ", an)
	newS = strconv.Itoa(an)

	fmt.Printf("new string %s\n", newS)

}