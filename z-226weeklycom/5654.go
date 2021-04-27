package main

import "fmt"

func main() {
	l, h := 5, 15
	fmt.Println(countBalls(l, h))
}

func countBalls(lowLimit int, highLimit int) int {
	ans := make([]int, 45)
	for i := highLimit; i >= lowLimit; i-- {
		sum := 0
		fmt.Println(i)
		j := i
		for j > 0 {
			sum += j % 10
			j /= 10
		}
		ans[sum]++
	}
	res := 0
	for i := 1; i < len(ans); i++ {
		if ans[i] > res {
			res = ans[i]
		}
	}
	fmt.Println(ans)
	return res
}
