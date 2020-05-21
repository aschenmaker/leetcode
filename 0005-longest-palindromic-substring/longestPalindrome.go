package main

import "fmt"

func main() {
	s := "abaabab"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		evenLeft, evenRight := expandAroundCenter(s, i, i)
		oddLeft, oddRight := expandAroundCenter(s, i, i+1)
		if evenRight-evenLeft > end-start {
			start, end = evenLeft, evenRight
		}
		if oddRight-oddLeft > end-start {
			start, end = oddLeft, oddRight
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
