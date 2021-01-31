package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 10, 11, 12}}
	fmt.Println(findDiagonalOrder(matrix))
}

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	row := len(matrix)
	col := len(matrix[0])

	if row == 1 {
		return matrix[0]
	}

	result := []int{}
	flag := 1

	for i := 0; i < row; i++ {
		cur := []int{}
		vi, j := i, 0
		for vi >= 0 && j < col {
			cur = append(cur, matrix[vi][j])
			vi--
			j++
		}
		if flag == 1 {
			// 类似于 js中es6的写法
			result = append(result, cur...)
		} else {
			cur = reverse(cur)
			result = append(result, cur...)
		}
		flag *= -1
	}

	for j := 1; j < col; j++ {
		cur := []int{}
		i, vj := row-1, j
		for vj < col && i >= 0 {
			cur = append(cur, matrix[i][vj])
			i--
			vj++
		}
		if flag == 1 {
			// 类似于 js中es6的写法
			result = append(result, cur...)
		} else {
			cur = reverse(cur)
			result = append(result, cur...)
		}
		flag *= -1
	}
	return result
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
