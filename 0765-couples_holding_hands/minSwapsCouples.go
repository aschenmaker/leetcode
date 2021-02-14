package main

import "fmt"

func main() {
	row := []int{1, 2, 0, 3}
	row2 := []int{1, 2, 0, 3}
	fmt.Println(minSwapsCouples(row))
	fmt.Println(useUnionFind(row2))
}

// 使用谈心算法，进行交换；
// 使用异或，寻找匹配对象
func minSwapsCouples(row []int) int {
	n := len(row)
	res := 0
	for i := 0; i < n-1; i += 2 {

		if row[i] == row[i+1]^1 {
			continue
		}
		for j := i + 1; j < n; j++ {
			if row[i] == row[j]^1 {
				row[i+1], row[j] = row[j], row[i+1]
			}

		}
		res++
	}
	return res
}

// 使用并查集
func useUnionFind(row []int) int {
	n := len(row)
	uf := newUnionFind(n / 2)
	for i := 0; i < n; i += 2 {
		uf.union(row[i]/2, row[i+1]/2)
	}
	return n/2 - uf.setCount
}
