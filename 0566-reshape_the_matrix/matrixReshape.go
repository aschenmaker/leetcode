package main

import "fmt"

func main() {
	nums := [][]int{{1, 2}, {3, 4}}
	r, c := 1, 4
	fmt.Println(matrixReshape(nums, r, c))
	fmt.Println(newmatrixReshape(nums, r, c))
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	curR := len(nums)
	curC := len(nums[0])
	if curC*curR != r*c {
		return nums
	}
	var num []int
	for i := 0; i < curR; i++ {
		for j := 0; j < curC; j++ {
			num = append(num, nums[i][j])
		}
	}

	ans := make([][]int, r)
	for i := range ans {
		ans[i] = make([]int, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			ans[i][j] = num[i+j]
		}
	}
	return ans
}

func newmatrixReshape(nums [][]int, r int, c int) [][]int {
	curR := len(nums)
	curC := len(nums[0])
	if curC*curR != r*c {
		return nums
	}

	ans := make([][]int, r)
	for i := range ans {
		ans[i] = make([]int, c)
	}

	for i := 0; i < curR*curC; i++ {
		ans[i/c][i%c] = nums[i/curC][i%curC]
	}
	return ans
}
