package main

import "fmt"

func main() {
	A := []int{4, 5, 0, -2, -3, 1}
	K := 5
	fmt.Println(subarraysDivByK(A, K))
}

func subarraysDivByK(A []int, K int) int {
	hashmap := map[int]int{0: 1}
	sum, ans := 0, 0
	for _, elem := range A {
		sum += elem
		modules := (sum%K + K) % K
		ans += hashmap[modules]
		hashmap[modules]++
	}
	return ans
}
