package main

import "fmt"

func main() {
	matrix := [][]int{{1, 0, 3}}
	fmt.Println(setZeroes(matrix))
}

// 使用第一行 第一列 来作为标记，对改行存在为0的数 进行标记 。
func setZeroes(matrix [][]int) [][]int {
	isFirstRowZero := false
	isFirstColZero := false

	rowLen := len(matrix)
	colLen := len(matrix[0])
	for i, row := range matrix {

		for j, elem := range row {

			if elem == 0 && i == 0 {
				isFirstRowZero = true
			}
			if elem == 0 && j == 0 {
				isFirstColZero = true
			}
			if elem == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < rowLen; i++ {
		row := matrix[i]
		if row[0] == 0 {
			for j := range row {
				matrix[i][j] = 0
			}
		}
	}

	for j := 1; j < colLen; j++ {
		elem := matrix[0][j]
		if elem == 0 {
			for i := range matrix {
				if i == 0 {
					continue
				}
				matrix[i][j] = 0
			}
		}
	}

	if isFirstRowZero {
		for i := 0; i < colLen; i++ {
			matrix[0][i] = 0
		}
	}

	if isFirstColZero {
		for i := 0; i < rowLen; i++ {
			matrix[i][0] = 0
		}
	}

	return matrix
}
