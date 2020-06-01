package main

import "fmt"

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	// 找到最大值
	n := len(candies)
	maxCandies := 0
	for i := 0; i < n; i++ {
		maxCandies = max(maxCandies, candies[i])
	}
	ret := make([]bool, n)
	for i := 0; i < n; i++ {
		ret[i] = candies[i]+extraCandies >= maxCandies

	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
