package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 1, 1}
	k := 3
	fmt.Println(numberOfSubarrays(nums, k))
}

func numberOfSubarrays(nums []int, k int) int {
	var cnt = make([]int, len(nums)+1)
	cnt[0] = 1
	sum, result := 0, 0
	for _, num := range nums {
		sum += num & 1 // 累加当前的奇数个数
		cnt[sum]++     // 累计函数++

		if sum >= k {
			result += cnt[sum-k]
		}
	}
	fmt.Println(cnt)
	return result
}
