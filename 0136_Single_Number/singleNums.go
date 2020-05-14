package main

import "fmt"

func singleNumber(nums []int) int {
	a := 0
	for _, v := range nums {
		a ^= v
	}
	return a
}

func main() {
	var nums = []int{1, 1, 2, 3, 2}
	ret := singleNumber(nums)
	fmt.Println(ret)
}
