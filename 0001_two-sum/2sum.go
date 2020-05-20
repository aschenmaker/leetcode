package main

import "fmt"

func main() {
	var nums = []int{2, 7, 11, 15, 5, 4}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	ret := []int{}
	m := make(map[int]int)
	for i, k := range nums {
		if value, exist := m[target-k]; exist {
			ret = append(ret, value)
			ret = append(ret, i)
		}
		m[k] = i
	}
	return ret
}
