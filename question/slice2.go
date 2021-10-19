package main

import (
	"fmt"
)

//func main()  {
//	s1 := [3]int{1,2,3}
//	s2 := s1[1:]
//	s2[1] = 4;
//	fmt.Println(s1)
//	s2 = append(s2, 5, 6, 7)
//	fmt.Println(s1)
//	fmt.Println(s2)
//
//	m := map[int]string{0 :"zero", 1:"one"}
//	for k, v := range m{
//		fmt.Println(k, v)
//	}
//}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) // 10, 1, 2, 3 // 1 1 3 4
	a = 0
	defer calc("2", a, calc("20", a, b)) // 20, 0, 2, 2 // 2, 0, 2, 2
	b = 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
