package main

import (
	"fmt"
	"math"
)

func main() {
	S := "ADOBECODEBANC"
	T := "ABC"
	ret := minWindow(S, T)
	fmt.Println(ret)
}

func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}
	fmt.Println(ori)
	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for left, right := 0, 0; right < sLen; right++ {
		if right < sLen && ori[s[right]] > 0 {
			cnt[s[right]]++
		}
		for check() && left <= right {
			if right-left+1 < len {
				len = right - left + 1
				ansL, ansR = left, left+len
			}
			if _, ok := ori[s[left]]; ok {
				cnt[s[left]]--
			}
			left++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
