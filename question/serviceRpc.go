package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

type Params struct {
	height int
	width int
}

type Rect struct {

}

func (r *Rect) Area(p Params, ret *int) error  {
	*ret = p.height * p.width
	return nil
}

func (r *Rect) Perimeter(p Params, ret *int) error  {
	*ret = (p.height + p.width) * 2
	return nil
}

func main()  {
	// 注册服务
	rect := new(Rect)
	rpc.Register(rect) // 注册一个rect服务
	rpc.HandleHTTP() // 服务处理绑定到http协议上
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("注册成功")
}
