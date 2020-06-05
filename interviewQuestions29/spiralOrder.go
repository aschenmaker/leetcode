package main

import "fmt"

func main() {
	matrix := make([][]int, 3)
	n := 1
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 4)
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = n
			n++
		}
	}
	fmt.Println(matrix)
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	// 初始化遍历过数组
	rows, columns := len(matrix), len(matrix[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	var (
		total           = rows * columns
		result          = make([]int, total)
		row, column     = 0, 0
		directions      = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
		directionsIndex = 0
	)
	for i := 0; i < total; i++ {
		result[i] = matrix[row][column]
		visited[row][column] = true
		nextRow, nextColumn := row+directions[directionsIndex][0], column+directions[directionsIndex][1]
		fmt.Println(nextRow, nextColumn)
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			directionsIndex = (directionsIndex + 1) % 4
		}
		row += directions[directionsIndex][0]
		column += directions[directionsIndex][1]
	}
	return result
}

// 使用边界✅
