package main

import (
	"fmt"
	"time"
)
var fibs [41]uint64
func main()  {
	var num int = 41
	println(fibs[20])
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < num; i++  {
		result = fbl(i)
		fmt.Printf("fbl(%d) is : %d\n", i, result)
	}

	for i := 0; i < num; i++  {
		result = fbl(i)
		//fmt.Printf("fbl(%d) is : %d\n", i, result)
	}
	end := time.Now()

	delta := end.Sub(start)
	fmt.Printf("用时%s\n", delta) // 未使用内存缓存：1.3893438s  使用内存缓存：3.9937ms
	println(fibs[20])

}

//获取斐波拉数列
func fbl(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if (n <= 1) {
		res = 1
	} else {
		res = fbl(n-1) + fbl(n-2)
	}

	fibs[n] = res
	return
}