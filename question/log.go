package main

import (
	"fmt"
	"log"
)

func main() {
	// log.Fatal("fatal level log: log entry") // 输出日志，终止程序
	log.Println("nomal level log")

	data := "A\xfe\x02\xff\x04"
	for _, v := range data {
		fmt.Printf("%#x ", v)
	}

	for _, v := range []byte(data) {
		fmt.Printf("%#x ", v)
	}

	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}

	isSpace := func(char byte) bool {
		switch char {
		case ' ':
			fallthrough //强制执行下一个case
		case 'a':
			fallthrough
		case '\t':
			return true
		}
		return false
	}

	fmt.Println(isSpace('\t'))
	fmt.Println(isSpace(' '))
	fmt.Println(isSpace('a'))

	datas := []int{1, 2, 3}

	i := 0

	i++ // go 只有i++， 没有++i
	fmt.Println(datas[i])

}
