package main

var a string

func main() {
	var b float64 = 8
	var c float64 = 0
	r := b / c
	print(r)
	a = "G"
	print(a) // G
	f1()  // 0 0
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}

