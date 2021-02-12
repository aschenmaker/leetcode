package main

import "fmt"

func main() {
	fmt.Println(getRow(10))
	fmt.Println(getRowBetter(10))
	fmt.Println(getRowByMath(10))
}

func getRow(rowIndex int) []int {
	c := make([][]int, rowIndex+1)
	for i := range c {
		c[i] = make([]int, i+1)
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}

	return c[rowIndex]
}

// 进一步优化，使用滚动数组来记录前一段的数组
func getRowBetter(rowIndex int) []int {
	var pre, cur []int
	for i := 0; i <= rowIndex; i++ {
		cur = make([]int, i+1)
		cur[0], cur[i] = 1, 1
		for j := 1; j < i; j++ {
			cur[j] = pre[j-1] + pre[j]
		}
		pre = cur
	}
	return pre
}

// 最后使用公式，
func getRowByMath(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}
	return row
}
