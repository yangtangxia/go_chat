package main

import (
	"fmt"
)

func main()  {
	for i:= 0; i <= 25; i++ {
		for j := 0; j <= i; j++ {
			print("G")
		}
		print("\n")
	}

	str := "A"
	for i := 1; i <= 25; i++ {
		println(str)
		str += "A"
	}

	for i := 0; i<= 10; i++ {
		fmt.Printf("%b is : %b \n", i, ^i)
	}

	for i := 1; i <=100; i++ {
		if (i % 3 == 0) && (i % 5 == 0)  {
			print("FizzBuzz")
		} else if i % 3 == 0 {
			print("Fizz")
		} else if i % 5 == 0 {
			print("Buzz")
		} else {
			print(i)
		}
		println()
	}

	for i := 1; i <= 10; i++ {
		for j :=1; j<= 20; j++ {
			print("*")
		}
		println()
	}
}
