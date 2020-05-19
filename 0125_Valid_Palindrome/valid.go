package main

import (
	"fmt"
)

func main() {
	s := " "
	ret := isPalindrome(s)

	fmt.Println(ret)
}

func isPalindrome(s string) bool {
	// s = strings.ToLower(s)
	x, y := 0, len(s)-1
	for x < y {
		if !(s[x] > 47 && s[x] < 58 || s[x] > 96 && s[x] < 123) {
			x++
		} else if !(s[y] > 47 && s[y] < 58 || s[y] > 96 && s[y] < 123) {
			y--
		} else if s[x] != s[y] {
			return false
		} else {
			x++
			y--
		}
	}
	return true
}
