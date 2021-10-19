package main

import (
	"fmt"
)

func main()  {
	k := 5
	switch k {
	case 4:
		fmt.Println("k <= 4")
		fallthrough
	case 5:
		fmt.Println("k <= 5")
		fallthrough
	case 6:
		fmt.Println("k<=6")
		fallthrough
	case 7:
		fmt.Println("k<=7")
		//fallthrough
	default:
		fmt.Println("default case")
	}

	fmt.Println("k value", k)
}
