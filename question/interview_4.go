package main

import "fmt"

var (
	size = 1024
	max_size = size * 2
)

func main()  {
	fmt.Println(size, max_size)

	sn1 := struct {
		age int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		age int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m map[string]string
	}{age: 11, m: map[string]string{"a":"1"}}

	sm2 := struct {
		age int
		m map[string]string
	}{age: 11, m: map[string]string{"a":"1"}}

	//只有bool、数值型、字符、指针、数组可以比较是否相等，切片、map、函数不能比较的； 结构体比较只有结构相同，结构值也相同、顺序也相同才相等
	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}
