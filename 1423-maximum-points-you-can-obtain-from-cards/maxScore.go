package main

import "fmt"

func main() {
	cardPoints := []int{1, 1000, 1}
	k := 2
	fmt.Println(maxScore(cardPoints, k))
}

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	left, right := n-k, n-1
	windowSum := 0
	for i := left; i <= right; i++ {
		windowSum += cardPoints[i]
	}

	ans := windowSum

	if k == n {
		return ans
	}

	for left <= n-1 {
		right++
		fmt.Println(cardPoints[left%n], cardPoints[right%n])
		windowSum = windowSum - cardPoints[left%n] + cardPoints[right%n]
		left++
		if windowSum > ans {
			ans = windowSum
		}
	}
	return ans
}
