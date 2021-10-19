package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main()  {

	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}

	where()

	where()

	where()

	println(2222)
	log.SetFlags(log.Llongfile)
	log.Print("")

	var wheres = log.Print

	wheres()

	// 函数执行时间\
	start := time.Now() // 当前时间
	sum(1000000)
	end := time.Now()
	delta := end.Sub(start)

	fmt.Printf("sum took time %s\n", delta)

}

func sum(n int) (sum int)  {
	for i :=0; i< n; i++ {
		sum += i
	}
	return
}

