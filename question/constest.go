package main

import (
	"fmt"
	"os"
	"runtime"
)

const Ln2 = 0.274329847324728483472384247239574302542350432
const Log2e = 1 / Ln2
const Billion =  1e9
const hardEight = (1 << 100) > 97

const (
	a = iota // 0
	f = iota // 1
	c = iota // 2
)

const (
	x = iota
	y
	z
	d = 5
	e
)

const (
	apple, banana = iota + 1, iota + 2
	cherimoya, durian // 1+ 1, 1 + 2
	elderberry, fig
	fx, fy
)
// 1  00000100
const (
	Open = 1 << iota
	Close
	Pending
)

func main()  {
	fmt.Println(Log2e)
	fmt.Println(Billion)
	fmt.Println(hardEight)

	fmt.Println(a, f, c) // 0 1 2\
	fmt.Println(x, y, z, d, e) // 0 1 2 5 5

	fmt.Println(apple, banana, cherimoya, durian)
	fmt.Println(apple, banana, cherimoya, durian, elderberry, fig)
	fmt.Println(apple, banana, cherimoya, durian, elderberry, fig, fx,fy)

	fmt.Println(Open, Close, Pending)

	var goos string = runtime.GOOS
	fmt.Printf("the operating system is: %s \n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("path is %s\n", path)
}
