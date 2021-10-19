package main

import (
	"fmt"
)

func main()  {
	seasons := []int{10, 20, 30, 40}
	for k, item := range seasons {
		seasons[k] = item * 2
	}

	fmt.Print(seasons)
	println()
	arrF := [3]float32{1, 2, 4.5}

	fmt.Printf("arrF sum is: %f\n", Sum(arrF))

	slic_arr := arrF[:2]
	fmt.Printf("slic_arr sum is : %f\n", slicSum(slic_arr))
	avg_arr := []int{2, 3, 4, 1}
	a := avg_arr[2:3]
	fmt.Printf("length is :%d, cap is %d\n", len(a), cap(a))
	sum, avg := SumAndAverage(avg_arr)
	fmt.Printf("sum : %d and avg : %f\n", sum, avg)

	fmt.Printf("avg_arr is minValue is : %d\n", minSlice(avg_arr))
	fmt.Printf("avg_arr is maxValue is : %d\n", maxSlice(avg_arr))
}

func Sum(arrF [3]float32) (sum float32)  {
	for _, value := range arrF {
		sum += value
	}
	return
}

func slicSum(arrF []float32) (sum float32)  {
	for _, value := range arrF {
		sum += value
	}
	return
}

func SumAndAverage(arr []int) (int, float32)  {
	sum := 0
	for _, value := range arr {
		sum += value
	}

	length := len(arr)
	avg := float32(sum / length)
	return sum, avg
}

func minSlice( slice []int) (min int)  {
	min = slice[0]
	for _, value := range slice {
		if min > value {
			min = value
		}
	}
	return
}

func maxSlice( slice []int) (max int) {
	max = slice[0]
	for _, value := range slice {
		if max < value {
			max = value
		}
	}
	return
}


