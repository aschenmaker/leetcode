package main

import "fmt"

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	left, right := 1, x
	for left < right {
		mid := (left + right + 1) >> 1

		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid
		}
	}

	return left
}

func main() {
	ret := mySqrt(4)
	fmt.Println(ret)
}
