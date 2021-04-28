package main

import (
	"fmt"
	"math"
)

func main() {
	a := 1
	fmt.Println(judgeSquareSum(a))
	b := 5
	fmt.Println(judgeSquareSum(b))
	c := 16
	fmt.Println(judgeSquareSum(c))
	d := 2147483644
	fmt.Println(judgeSquareSum(d))
}

func judgeSquareSum(c int) bool {
	if c == 0 || c == 1 {
		return true
	}
	for a := 0; a*a < c; a++ {
		rt := math.Sqrt(float64(c - a*a))
		if rt == math.Floor(rt) {
			return true
		}
	}
	return false
}

func judgeSquareSum2(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		}
		if sum > c {
			right--
		}
		if sum < c {
			left++
		}
	}
	return false
}
