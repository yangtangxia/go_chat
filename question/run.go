package main

/**
随机彩票数
 */
import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main()  {
	getBalls()
}

func getBalls()  {
	var redBalls [6]int
	for i := 0; i < 6; i++ {
		for  {
			rand.Seed(time.Now().UnixNano())
			num := rand.Intn(33) + 1
			if ifInBalls(num, redBalls) {
				redBalls[i] = num
				break
			} else {
				continue
			}
		}
	}

	sort.Ints(redBalls[:])

	var blueBall [1] int
	num := rand.Intn(16) + 1
	blueBall[0] = num

	fmt.Println(redBalls, blueBall)
}

func ifInBalls(param int, list [6]int) bool  {
	for _, b := range list {
		if b == param {
			return false
		}
	}
	return true
}