package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
	fmt.Println(productExceptSelfO1(nums))
}

func productExceptSelf(nums []int) []int {
	length := len(nums)
	L, R, answers := make([]int, length), make([]int, length), make([]int, length)

	L[0] = 1
	for i := 1; i < length; i++ {
		L[i] = nums[i-1] * L[i-1]
	}
	R[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		R[i] = nums[i+1] * R[i+1]
	}

	for i := 0; i < length; i++ {
		answers[i] = L[i] * R[i]
	}
	return answers
}

// o(1)
func productExceptSelfO1(nums []int) []int {
	length := len(nums)
	answers := make([]int, length)

	answers[0] = 1
	for i := 1; i < length; i++ {
		answers[i] = nums[i-1] * answers[i-1]
	}

	R := 1
	for i := length - 1; i >= 0; i-- {
		answers[i] *= R

		R *= nums[i]
	}
	return answers
}
