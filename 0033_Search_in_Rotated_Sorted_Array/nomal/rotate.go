package main

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if target == nums[0] {
		return 0
	}
	if target == nums[n-1] {
		return n - 1
	}
	if target > nums[0] {
		i := 0
		for i+1 < n-1 && nums[i] < nums[i+1] {
			if nums[i+1] == target {
				return i + 1
			}
			i++
		}
		return -1
	}
	if target < nums[0] {
		i := n - 1
		for i-1 > 0 && nums[i-1] < nums[i] {
			if target == nums[i-1] {
				return i - 1
			}
			i--
		}
		return -1
	}
	return -1
}

func main() {
	var nums = []int{4, 5, 6, 7, 0, 1, 2}
	target := 3

	fmt.Println(search(nums, target))
}
