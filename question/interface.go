package main

import "fmt"

//type Cat struct {
//
//}
//
//type Dog struct {
//
//}

//func (c Cat) Say() string  {
//	return "喵喵喵"
//}
//
//func (d Dog) Say() string  {
//	return "汪汪汪"
//}

//func main()  {
//	c := Cat{}
//	fmt.Println("喵", c.Say())
//
//	d := Dog{}
//	fmt.Println("汪", d.Say())
//}

type Sayer interface {
	say()
}

type dog struct {
}

type cat struct {
}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func main() {
	var x Sayer
	a := cat{}
	b := dog{}
	x = a
	x.say()
	x = b
	x.say()

}
