package main

import (
	"fmt"
)

func main()  {

	//b()
	//错误语法
	//var m map[string]int
	//m["one"] = 1
	//正确语法
	//var s []int
	//s = append(s, 1)
	//fmt.Println(s)

	//判断key是否存在
	x := map[string]string {"one":"2", "two": "", "three": "3"}
	//错误示例
	if v := x["two"]; v== "" {
		fmt.Println("key two is no entry")
	}
	//正确示例
	if _, ok := x["twos"]; !ok {
		fmt.Println("key two is no entry")
	}

	// 修改字符串
	s := "text杨"
	//x[0] = "T" // 错误
	sBytes := []byte(s) // 将 string 转为 []byte 修改后，再转为 string 即可。
	fmt.Println(sBytes)
	sBytes[0] = 'T'
	s = string(sBytes)
	fmt.Println(s)

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for key, val := range slice {
		value := val
		m[key] = &value
	}
	fmt.Println(m)
	for k, v := range m{
		fmt.Println(k, "=>", *v)
	}


}



//先进后出
func b()  {
	for i := 0; i< 10; i++ {
		defer fmt.Print(i)
	}
}