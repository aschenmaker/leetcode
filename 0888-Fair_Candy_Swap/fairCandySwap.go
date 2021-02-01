package main

import "fmt"

func main() {
	A := []int{2}
	B := []int{1, 3}
	fmt.Println(fairCandySwap(A, B))
	fmt.Println(fairCandySwapBetter(A, B))
}

func fairCandySwap(A []int, B []int) []int {
	sumA, sumB := 0, 0
	for _, v := range A {
		sumA += v
	}
	for _, v := range B {
		sumB += v
	}
	if sumA > sumB {
		dif := (sumA - sumB) / 2
		for i := 0; i < len(A); i++ {
			for j := 0; j < len(B); j++ {
				if A[i]-B[j] == dif {
					return []int{A[i], B[j]}
				}
			}
		}
	} else {
		dif := (sumB - sumA) / 2
		for i := 0; i < len(A); i++ {
			for j := 0; j < len(B); j++ {
				if B[j]-A[i] == dif {
					return []int{A[i], B[j]}
				}
			}
		}
	}
	return []int{}
}

// 优化 使用map结构 来进行两数之和 M+N 的时间复杂度

func fairCandySwapBetter(A []int, B []int) []int {
	sumA, sumB := 0, 0
	setA := map[int]struct{}{}

	for _, v := range A {
		sumA += v
		setA[v] = struct{}{}
	}
	for _, v := range B {
		sumB += v
	}

	dif := (sumA - sumB) / 2

	for i := 0; ; i++ {
		y := B[i]
		x := y + dif
		if _, has := setA[x]; has {
			return []int{x, y}
		}
	}

}
