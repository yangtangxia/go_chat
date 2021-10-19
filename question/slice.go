package main

import "fmt"

var arrs = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arrs[1:3]
var slice1 []int = arrs[:5]
var slice2 []int = arrs[2:]
var slice3 []int = arrs[:]
var slice4 []int = arrs[:len(arrs)-1]

func main() {

	s := make([]byte, 5)
	s = s[2:4]
	print(len(s))
	println()
	print(cap(s))
	println()
	//b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	b := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Print(b[1:4])
	println()
	fmt.Print(b[:2])
	println()
	fmt.Print(b[2:])
	println()
	fmt.Print(b[:])
	println()
	b2 := b[2:]
	fmt.Print(b2)
	b2[1] = "j"
	fmt.Print(b2)
	println()
	fmt.Print(b)
	//var s1 []int
	//fmt.Println(s1)
	//if s1 == nil {
	//	fmt.Println("空的")
	//} else {
	//	fmt.Println("不是空的")
	//}
	//
	//s2 := []int{}
	//fmt.Println(s2)
	//var s3 []int = make([]int, 0)
	//fmt.Println(s3)
	//var s4 []int = make([]int, 0, 0)
	//fmt.Println(s4)
	//s5 := []int{1, 2, 3}
	//fmt.Println(s5)
	//
	//s6 := [5]int{1, 2, 3,4, 5}
	//var s7 []int
	//s7 = s6[1:4]
	//fmt.Println(s7)

	//fmt.Println(arrs, slice0, slice1, slice2, slice3, slice4)
	//
	//fmt.Println(arrs[2])

	//data := [...]int{0, 1, 2, 3, 4, 5}
	//s := arrs[2: 4]
	//s[0] += 100
	//s[1] += 200
	//fmt.Println(s)
	//fmt.Println(arrs)
	//
	//s1 := []int{0, 1, 2, 3, 8:100}
	//fmt.Println(s1, len(s1), cap(s1))
	//
	//s2 := make([]int, 6, 8)
	//fmt.Println(s2, len(s2), cap(s2))
	//s2[5] = 100
	//fmt.Println(s2)
	//var s3 = []int{1, 2, 3, 5, 6, 7}
	//s4 := append(s2, s3... )
	//s4 = append(s4, 200, 300, 400)
	//fmt.Println(s4)

	//x := []int{1, 2, 3}
	//func(arr []int) {
	//	arr[0] = 7
	//	fmt.Println(arr)
	//}(x)
	//fmt.Println(x)
	//
	//// cap 分配
	//slice := make([]int, 0, 1)
	//c := cap(slice)
	//
	//fmt.Println(c)
	//for i := 0; i < 50; i++ {
	//	slice = append(slice, i)
	//	if n := cap(slice); n > c {
	//		fmt.Printf("cap: %d -> %d\n", c, n)
	//		c = n
	//	}
	//}
	//
	//fmt.Println(slice)
	//// 以逗号分隔切片为字符串
	//ss := strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", ",", -1)
	//fmt.Println(ss)
	//// 分隔数组为字符串
	//array1 := [6]int{1, 2, 3, 4, 5}
	//array_string := strings.Replace(strings.Trim(fmt.Sprint(array1), "[]"), " ", ",", -1)
	//fmt.Println(array_string)
	//
	//var num int
	//fmt.Println(&num)
	//
	//ptr := &num
	//*ptr = 20
	//fmt.Println(num)

}
