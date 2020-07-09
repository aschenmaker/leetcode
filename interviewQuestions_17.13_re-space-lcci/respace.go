package main

import "fmt"

func main() {
	dictionary := []string{"looked", "just", "like", "her", "brother"}
	sentence := "jesslookedjustliketimherbrother"

	fmt.Println(respace(dictionary, sentence))
}

func respace(dictionary []string, sentence string) int {
	n, inf := len(sentence), 0x3f3f3f3f3f
	root := &Trie{next: [26]*Trie{}}
	for _, word := range dictionary {
		root.insert(word)
	}

	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = inf
	}
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		curPos := root
		for j := i; j >= 1; j-- {
			t := int(sentence[j-1] - 'a')
			if curPos.next[t] == nil {
				break
			} else if curPos.next[t].isEnd {
				dp[i] = min(dp[i], dp[j-1])
			}
			if dp[i] == 0 {
				break
			}
			curPos = curPos.next[t]
		}
	}
	fmt.Println(inf)
	return dp[n]
}

// Trie is dictionary of words
type Trie struct {
	next  [26]*Trie
	isEnd bool
}

func (t *Trie) insert(s string) {
	curPos := t
	for i := len(s) - 1; i >= 0; i-- {
		temp := int(s[i] - 'a')
		if curPos.next[temp] == nil {
			curPos.next[temp] = &Trie{next: [26]*Trie{}}
		}
		curPos = curPos.next[temp]
	}
	curPos.isEnd = true
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
