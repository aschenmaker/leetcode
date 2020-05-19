package main

import "fmt"

func main() {
	s := "aba"
	ret := validPalindrome(s)
	fmt.Println(ret)
}

func isPalindromeeRange(s string, i, j int) bool {
	x, y := i, j
	for x < y {
		if s[x] != s[y] {
			return false
		}
		x++
		y--
	}
	return true
}

func validPalindrome(s string) bool {
	x, y := 0, len(s)-1

	for x < y {
		if s[x] != s[y] {
			return (isPalindromeeRange(s, x+1, y) || isPalindromeeRange(s, x, y-1))
		}
		x++
		y--
	}

	return true
}
