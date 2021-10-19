package main

import "fmt"

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["梅西"] = 100
	scoreMap["内马尔"] = 90
	fmt.Println(scoreMap)
	fmt.Printf("type of a:%T\n", scoreMap)
}
