package main

import (
	"fmt"
)

func main()  {
	//x := min(1, 3, 2, 0, -8)

	//fmt.Printf("the minium is: %d \n", x)
	//slice := []int{7, 9, 3, 5, 1}
	//x = min(slice ...)
	//fmt.Printf("the minimum in the slice is :%d\n", x)

	//printlns(slice...)

	//functionDefer() // defer

	//i := 0
	//defer fmt.Println(i)
	//return
	//a()
	res := 0
	res = fibonacci(2)
	println(res)
	ress := fbn(10)
	fmt.Println(ress)

	fmt.Printf("%d is even : is %t\n", 16, even(16))
	fmt.Printf("%d is odd : is %t\n", 17, odd(17))

	fmt.Printf("%d is odd: is %t\n", 18, odd(18))

	//dayin(10)
	//jiec(5)
	fmt.Printf("阶乘 ： %d\n", jiec(38))

	fv := func() { fmt.Printf("Hello World")}
	//fv()
	fmt.Printf("fv type is %T, msg is%v\n", fv, fv)
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}

	min := s[0]
	for _, v := range s{
		if v < min {
			min = v
		}
	}
	return min
}

// 变成函数，对每个元素换行打印
func printlns(s ...int)  {
	if len(s) == 0 {
		fmt.Println("请传入不为空数据")
	}

	for _, v := range s{
		println(v)
	}
}

func functionDefer()  {
	defer function1()
	fmt.Printf("in function at th top \n")
	defer function2()
	fmt.Printf("in function 2 at the bottom\n")
	defer function3()
}
func function1()  {
	fmt.Printf("function1:deferred one\n")
}
func function2()  {
	fmt.Printf("function2:deferred two\n")
}
func function3()  {
	fmt.Printf("function3:deferred three\n")
}

func a()  {
	for i := 0; i< 5; i++ {
		defer println(i)
	}
}

//获取斐波拉数列
func fibonacci(n int) int  {
	if (n <= 1) {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

//获取斐波拉数
func fbn(n int) []int  {
	var slice = make([]int, n)
	//slice[1] = 1
	//slice[2] = 1
	//if n > 2 {
		for i := 0; i < n; i++ {
			//slice[i] = slice[i - 1] + slice[i - 2]
			slice[i] = fibonacci(i)
		}
	//}
	return slice
}

func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(recsign(nr) - 1)
}

func odd(nr int) bool  {
	if nr == 0 {
		return false
	}
	return even(recsign(nr) - 1)
}

func recsign(nr int) int  {
	if nr < 0 {
		return -nr
	}
	return nr
}


//递归打印10 到 1
func dayin(n int) int  {
	res := 0
	if n == 1 {
		res = 1
		println(res)

	} else {
		println(n)
		res = dayin(n - 1)
	}
	return res
}

//计算数字阶乘
func jiec(n int) int {
	res := 1
	if n == 1 {
		res = 1
	} else {
		res = jiec(n - 1) * n
	}

	return res
}
