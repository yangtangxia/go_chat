package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

//全局定义
var arr [2]int
var arr1 = [5]int{1, 2, 3, 4,5}
var arr2 = [3]int{2:2, 1:1}
var arr3 = [...]int{} // 定义空数组
var arr4 = [0]int{} // 空数组

var arr5 [5][3]int
var arr6 [2][4]int = [...][4]int{{1, 2,3},{2,3,4}}

func main()  {

	var arry1 = new([5]int)
	var arry2 [5]int
	fmt.Printf("arry1 type %T\n", arry1)
	fmt.Printf("arryt2 type %T\n", arry2)
	arry3 := *arry1
	arry3[2] = 80
	arry4 := arry2
	arry4[2] = 90
	fmt.Print(arry1)
	fmt.Print(arry3)
	fmt.Print(arry2)
	fmt.Print(arry4)

	//数组指针传递函数
	array1 := [3]int{8, 7, 9}
	result := sumTotal(&array1)
	println(result)
	os.Exit(20)
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(len(arr3))
	fmt.Println(arr4)

	//
	d := [...]struct{
		name string
		age uint8
	} {
		{"梅西", 34},
		{"内马尔", 29}, // 最后一行逗号必须
	}

	fmt.Println(d)

	// 多维数组
	fmt.Println(arr5)
	fmt.Println(arr6)

	a := [2]int{}
	fmt.Printf("a: %p\n", &a)
	test(a)
	fmt.Println(a)

	for k1, v1 := range arr6 {
		for k2, v2 := range v1 {
			fmt.Printf("(%d, %d) = %d", k1, k2, v2)
		}
	}

	rand.Seed(time.Now().Unix())
	bs := rand.Intn(100)
	fmt.Println(bs)

	ab := [5]int{1,3,5,7,8}
	checkValue(ab, 8)

}

func test(x [2]int)  {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}

func checkValue(a [5]int, target int)  {
	for i := 0; i < len(a); i++ {
		for j := 0; j <len(a); j++ {
			if (a[i] + a[j] == target) {
				fmt.Printf("(%d, %d)", i, j)
			}
		}
	}
}

func sumTotal(a *[3]int) (total int) {
	for _, v := range a {
		total += v
	}
	return
}