package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16))

	fmt.Println(isPerfectSquare(15))

	fmt.Println(isPerfectSquareBy(256))
}

func isPerfectSquare(num int) bool {
	i := 1
	for num > 0 {
		num -= i
		i += 2
	}
	if num == 0 {
		return true
	}
	return false
}

func isPerfectSquareBy(num int) bool {
	if num < 2 {
		return true
	}
	min, max := 0, num/2+1

	for max >= min {
		mid := (min + max) / 2
		if num == mid*mid {
			return true
		} else if num > mid*mid {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return false
}
