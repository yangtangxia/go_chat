package main

import (
	"fmt"
)

func main()  {
	slfrom := []int{1, 2, 3}
	slto := make([]int, 3)
	fmt.Printf("length:%d, cap : %d\n", len(slto), cap(slto))
	n := copy(slto, slfrom)
	fmt.Println(slto)
	fmt.Println(n)
	slic := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	slto = append(slto, slic...)
	fmt.Println(slto)
	fmt.Printf("length:%d, cap : %d\n", len(slto), cap(slto))

	s := make([]int, 3)
	fmt.Println(len(s))
	lenth := len(s)* 3
	cha := lenth - len(s)
	new_slic := make([]int, cha)
	new_s := append(s, new_slic...)
	fmt.Println(new_s)

	s1 := []int{2, 3}
	s2 := []int{1, 5, 9, 7}
	result := insertStringSlice(s1, s2, 1)
	fmt.Println(result)

	str := "abcdefg"
	s111 := str[(len(str)/2):] + str[:len(str)/2]

	fmt.Println(s111)
	row, res := splicString(str, 2)
	fmt.Println(row)
	fmt.Println(res)

	//reversal(str)
}

func insertStringSlice(s1 []int, s2 []int, n int) (result []int)  {
	if len(s2) < n {
		result = append(s2, s1...)
	} else {
		s2_1 := s2[:n]
		fmt.Println(s2_1) // 1
		s2_2 := s2[n:]
		fmt.Println(s2_2) // 5 9 7
 		s3 := append(s2_1, s1...)
		fmt.Println(s3) // 1 2 3
		fmt.Println(s2)
		result = append(s3, s2_2...)
		fmt.Println(result) // 1 2 3 2 3 7
		result = append(s2[:n], append(s1, s2[n:]...)...)
	}
	return
}
//要求其接受两个参数，原始字符串 str 和分割索引 i，然后返回两个分割后的字符串。
func splicString(str string, i int) (r1, r2 string) {
	r1 = str[:i]
	r2 = str[i:]
	return
}

func reversal(str string)  {
	s := []byte(str)
	new_s := []byte{}

	for i := len(s); i> 0; i-- {
		new_s = append(new_s, str[i])
	}

	fmt.Println(new_s)
}

type fun func()
func mapFunc(a fun)  {

}

func F()  {

}


