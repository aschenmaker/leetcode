package permutation

import (
	"sort"
)

func permutation(s string) []string {
	t := []byte(s)
	// 从小到大排序
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })

	n := len(t)
	perm := make([]byte, 0, n)
	vis := make([]bool, n)
	ans := []string{}

	var backtrace func(i int)
	backtrace = func(i int) {
		if i == n {
			ans = append(ans, string(perm))
			return
		}
		for k, b := range vis {
			// 确保从左往右进行遍历
			if b || k > 0 && !vis[k-1] && t[k-1] == t[k] {
				continue
			}
			vis[k] = true
			perm = append(perm, t[k])
			backtrace(i + 1)
			perm = perm[:len(perm)-1]
			vis[k] = false
		}

	}
	backtrace(0)

	return ans
}
