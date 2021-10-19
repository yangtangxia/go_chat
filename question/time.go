package main

import (
	"fmt"
	"time"
)

func main()  {
	timeDemo()
	timestampDemo()
}

func timeDemo()  {
	now := time.Now()
	fmt.Println("当前时间：time：%v\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day,hour, minute, second)
}

func timestampDemo()  {
	now := time.Now()
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Printf("current %v\n", timestamp1)
	fmt.Printf("current %v\n", timestamp2)
	timeobj := time.Unix(timestamp1, 0)
	fmt.Println(timeobj)
	year := timeobj.Year()
	fmt.Printf("%d", year)

	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}
}