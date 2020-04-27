package main

import "fmt"

func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	if l == 1 && nums[0] == target {
		return 0
	}
	left := 0
	right := l - 1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		if nums[mid] < nums[right] {
			if target <= nums[right] && target >= nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

func main() {
	var nums = []int{1, 2, 3}
	target := 3

	fmt.Println(search(nums, target))
}
