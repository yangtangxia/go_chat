package main

import (
	"fmt"
	"strings"
)

func main()  {
	p2 := Add2()
	fmt.Printf("call Add2 for 3 give : %v\n", p2(3))

	p3 := Adder(2)
	fmt.Printf("result : %v\n", p3(6))

	var f = adders()
	fmt.Print(f(1), "=") // 1
	fmt.Print(f(100), "=")
	fmt.Print(f(300), "=")

	var g int
	go func(i int) {
		s := 0
		for j := 0; j < i; j++ {
			s += j
		}
		g = s
		//println(g)
	}(1000)

	println(g)

	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")

	bmp := addBmp("file")
	jpeg := addJpeg("file")
	println(bmp)
	println(jpeg)
}

func Add2() func(b int) int  {
	return func(b int) int {
		return  b + 2
	}
}

func Adder(a int) func(b int) int  {
	return func(b int) int {
		return a + b
	}
}

func adders() func(int) int  {
	var x int
	println(x)
	return func(i int) int {
		x += i
		return x
	}
}

//动态追加后缀
func MakeAddSuffix(suffix string) func(string) string  {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}