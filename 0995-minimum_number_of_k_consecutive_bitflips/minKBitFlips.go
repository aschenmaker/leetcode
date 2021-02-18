package main

import "fmt"

func main() {
	A := []int{0, 0, 0, 1, 0, 1, 1, 0}
	k := 3
	fmt.Println(minKBitFlips(A, k))
	fmt.Println(useDiff(A, k))

}

// 逐一反转，最后校验
func minKBitFlips(A []int, K int) int {
	n := len(A)
	ans := -1
	flipTimes := 0
	for i, v := range A {
		if v == 0 && i+K <= n {
			for j := i; j < i+K; j++ {
				A[j] ^= 1
			}
			flipTimes++
			// fmt.Println(A)
		}
	}
	for i := n - K; i < n; i++ {
		if A[i] == 0 {
			return -1
		}
	}
	ans = flipTimes
	return ans
}

// 更优雅的方法，使用查分数组
func useDiff(A []int, K int) int {
	var ans int
	n := len(A)
	diff := make([]int, n+1)
	fmt.Println(diff)
	revCnt := 0
	for i, v := range A {
		revCnt ^= diff[i]
		if v == revCnt {
			if i+K > n {
				return -1
			}
			ans++
			revCnt ^= 1
			diff[i+K] ^= 1
		}
	}
	return ans
}
