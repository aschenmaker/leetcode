package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	length := len(nums)
	res := [][]int{}
	path := []int{}
	used := make([]int, length)

	for i := 0; i < length; i++ {
		used = append(used, 0)
	}
	res = dfs(nums, length, path, used, res)

	return res
}

func dfs(nums []int, length int, path []int, used []int, res [][]int) [][]int {
	// fmt.Println(path, res)
	if length == len(path) {
		res = append(res, path)
		// fmt.Println(res)
	} else {
		for i := 0; i < length; i++ {
			if used[i] == 0 {

				used[i] = 1
				// fmt.Println(res, path)
				res = dfs(nums, length, append(path, nums[i]), used, res)
				// pathlen := len(path)
				// path = path[0 : pathlen-1]
				used[i] = 0
			}
		}
	}
	return res
}

func main() {
	var nums = []int{1, 2, 3, 4}
	res := permute(nums)
	fmt.Println(res)
}
