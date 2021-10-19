package main

import (
	"fmt"
	"math"
	"math/big"
)

func main()  {

	im := big.NewInt(math.MaxInt64)
	in := im
	fmt.Println(im)
	fmt.Println(in)

	io := big.NewInt(3)
	fmt.Println(io)
	ip := big.NewInt(3)
	fmt.Printf("big int %v\n", ip)

	////ip.Mul(im, in).Add(ip, im).Div(ip, io)
	//io.Mul(ip, io) // 6 ip * io
	//fmt.Printf("io int %v\n", io)
	//io.Add(ip, io) // 8 ip + io
	//fmt.Printf("new io int %v\n", io)
	//io.Mul(ip,io).Add(ip, io)
	//fmt.Printf("new 222 io int %v\n", io)
	//x := big.NewInt(17)
	//io.Div(x, io) //  x/io 取整
	//fmt.Printf("new 1 int %v\n", io)

	io.Mul(ip, io).Add(ip, io).Div(ip, io)
	fmt.Printf("test %v\n", io)

	rm := big.NewRat(math.MaxInt64, 1956)
	fmt.Println(rm)
	rn := big.NewRat(-1956, math.MaxInt64)
	fmt.Println(rn)
	rp := big.NewRat(1, 1)
	fmt.Println(rp)
}
