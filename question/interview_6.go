package main

import "fmt"

type myint1 int

type myint2 = int

func main()  {
	var i int = 0
	var i1 myint1 = 0
	var i2 myint2 = i
	fmt.Println(i2)
	fmt.Println(i1)
}
