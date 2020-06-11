package main

import "fmt"

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	// temperatures:= []int{30,233}
	fmt.Println(dailyTemperatures(temperatures))
	fmt.Println(dailyTemperaturesStack(temperatures))
}

// 暴力

func dailyTemperatures(T []int) []int {
	length := len(T)
	ans := make([]int, length)
	for i := 0; i < length-1; i++ {
		var shift int = 1
		for j := i + 1; j < length; j++ {
			if T[i] < T[j] {
				ans[i] = shift
				break
			} else {
				shift++
			}
		}
	}
	return ans
}

// 单调栈
func dailyTemperaturesStack(T []int) []int {
	length := len(T)
	ans := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		temperature := T[i]
		for len(stack) > 0 && temperature > T[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[preIndex] = i - preIndex
		}
		stack = append(stack, i)
	}
	return ans
}
