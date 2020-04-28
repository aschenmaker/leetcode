package main

import "fmt"

func singleNumbers(nums []int) []int {
	var ret int = 0
	for i := range nums {
		ret = nums[i] ^ ret
	}
	bit := 1
	for i := 0; i < 32; i++ {
		if bit&ret != 0 {
			break
		}
		bit = bit << 1
	}

	var a int = 0
	var b int = 0
	for i := range nums {
		if nums[i]&bit != 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}

	// fmt.Println(a, b)
	return []int{a, b}
}

func main() {
	var nums = []int{4, 1, 4, 6, 8, 8, 6, 5}
	fmt.Println(singleNumbers(nums))
}
