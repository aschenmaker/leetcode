package main

import "fmt"

func main() {
	s := "ABAB"
	k := 2
	fmt.Println(characterReplacement(s, k))
}

// 滑动窗口
func characterReplacement(s string, k int) int {
	left, right, maxCnt := 0, 0, 0
	var mp [26]int
	for ; right < len(s); right++ {
		mp[s[right]-'A']++
		maxCnt = max(maxCnt, mp[s[right]-'A'])
		if right-left+1-maxCnt > k {
			mp[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
