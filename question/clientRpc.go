package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//type Params struct {
//	height, width int
//}

func main()  {
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Panicln(err)
	}

	ret := 0
	err2 := conn.Call("Rect.Area", Params{50, 60}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("面积：", ret)

	err3 := conn.Call("Rect.Perimeter", Params{50, 60}, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("周长：", ret)
}
