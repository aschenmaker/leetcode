package main

import "fmt"

func main() {
	n := 10
	fmt.Println(sumNums(n))
}

func sumNums(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}
