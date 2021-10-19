package main

import "fmt"

func main()  {
	var i1 = 5
	fmt.Printf("an integer : %d, pointer %s \n", i1, &i1)

	var intP *int

	intP = &i1
	i1 = 4
	fmt.Printf("point %s \n", *intP)

	var a = 5
	if (a > 10 ) {
		fmt.Print(1)
	} else {
		fmt.Print(2)
	}
}