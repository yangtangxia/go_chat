package main

import "fmt"

func main()  {
	LAB:
	for i := 0; i < 5; i++ {
		for j:= 0; j<= 5; j++ {
			if j == 4 {
				continue LAB // 退回到外层循环
				//break LAB  // 退出外层循环
				//break // 退出内层循环
			}

			fmt.Printf("i is: %d, and j is : %d\n", i, j)
		}
	}

	//i := 0
	//HERE:
	//	print(i)
	//	i++
	//	if i == 5 {
	//		return
	//	}
	//	goto HERE

	j := 0
	for  {
		if j >= 5 {
			break
		}
		fmt.Println("value is ", j)
		j++
	}
	fmt.Println("a atatus")
}
