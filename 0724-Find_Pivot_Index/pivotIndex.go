package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
}

func pivotIndex(nums []int) int {
	curSum := 0
	sum := 0
	for _, num := range nums {
		sum += num
	}
	for i := 0; i < len(nums); i++ {
		if sum-nums[i] == curSum*2 {
			return i
		}
		curSum += nums[i]
	}
	return -1
}
