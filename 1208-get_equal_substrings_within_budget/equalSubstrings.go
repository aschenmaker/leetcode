package main

import (
	"fmt"
	"math"
)

func main() {
	s := "abcd"
	t := "bcdf"
	fmt.Println(equalSubstring(s, t, 3))
}

func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	diff := make([]int, n)
	for i, ch := range s {
		diff[i] = int(math.Abs(float64(ch) - float64(t[i])))
	}
	maxLen := 0
	sum, start := 0, 0
	for end, d := range diff {
		sum += d
		for sum > maxCost {
			sum -= diff[start]
			start++
		}
		maxLen = max(maxLen, end-start+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
