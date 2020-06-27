package main

import "fmt"

func main() {
	var nums = []int{3, 5, -1, 1}
	fmt.Println(firstMissingPositive1(nums))
	fmt.Println(firstMissingPositive(nums))
}

// 声明了新的数组 占用的空间不满足题设
func firstMissingPositive1(nums []int) int {
	n := len(nums)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		if nums[i] > 0 && nums[i] < n {
			arr[nums[i]] = nums[i]
		}
	}
	max := 0
	for i := 1; i < len(arr); i++ {
		max = arr[i]
		if arr[i] != i {
			return i
		}
	}
	return max + 1
}

// 满足题设对复杂度的要求
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	fmt.Println(nums)
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
