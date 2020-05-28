package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "11[3[a]2[b0[c]]]"
	ret := decodeString(s)
	fmt.Println(ret)
}

func decodeString(s string) string {
	stack := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stack = append(stack, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stack = append(stack, string(cur))
			ptr++
		} else {
			ptr++
			sub := []string{}
			// pop the items before "["
			for stack[len(stack)-1] != "[" {
				sub = append(sub, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// exchange order of sub stack
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			// pop "["
			stack = stack[:len(stack)-1]
			// exchange the string type num to int time for calculating, than get num
			repTime, _ := strconv.Atoi(stack[len(stack)-1])
			// pop num
			stack = stack[:len(stack)-1]
			t := strings.Repeat(getString(sub), repTime)
			stack = append(stack, t)
		}

	}
	return getString(stack)
}

// concatnates arrays into string
func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}
