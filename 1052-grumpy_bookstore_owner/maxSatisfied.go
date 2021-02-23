package main

import (
	"fmt"
)

func main() {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	X := 3
	fmt.Println(maxSatisfied(customers, grumpy, X))
}

func maxSatisfied(customers []int, grumpy []int, X int) int {
	total := 0
	n := len(customers)
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			total += customers[i]
		}
	}
	increase := 0
	for i := 0; i < X; i++ {
		increase += customers[i] * grumpy[i]
	}
	maxIncrease := increase

	for i := X; i < n; i++ {
		increase = increase - customers[i-X]*grumpy[i-X] + customers[i]*grumpy[i]
		maxIncrease = max(increase, maxIncrease)
	}
	return total + maxIncrease
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
