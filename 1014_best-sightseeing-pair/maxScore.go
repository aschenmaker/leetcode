package main

import "fmt"

func main() {
	A := []int{8, 1, 5, 2, 6}
	fmt.Println(maxScoreSightseeingPair(A))
}

func maxScoreSightseeingPair(A []int) int {
	ans, mx := 0, A[0]+0
	for j := 1; j < len(A); j++ {
		ans = max(ans, mx+A[j]-j)
		mx = max(mx, A[j]+j)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
