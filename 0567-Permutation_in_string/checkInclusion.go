package main

import "fmt"

func main() {
	s1 := "ab"
	s2 := "caewdbas"
	fmt.Println(checkInclusion(s1, s2))
}

func checkInclusion(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	var count1, count2 [26]int
	for i, v := range s1 {
		count1[v-'a']++
		count2[s2[i]-'a']++
	}
	if count1 == count2 {
		return true
	}

	for i := n; i < m; i++ {
		count2[s2[i]-'a']++
		count2[s2[i-n]-'a']--
		if count1 == count2 {
			return true
		}
	}

	return false
}
