package main

import "fmt"

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}

// 使用双指针，分别在两个字符串中移动，最终判断 两个值 是否相等
func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == n
}
