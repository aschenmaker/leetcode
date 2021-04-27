package main

// func main() {
// 	s := "abcdefg"
// 	fmt.Println(checkPartitioning(s))
// }

func checkPartitioning(s string) bool {
	if len(s) < 3 {
		return false
	}
	for i := 1; i < len(s)-1; i++ {
		s1 := s[0:i]
		if !isOk(s1) {
			continue
		}
		for j := len(s) - 1; j > i; j-- {

			s3 := s[j:len(s)]

			if !isOk(s3) {
				continue
			}
			s2 := s[i:j]
			if isOk(s2) {
				return true
			}
			println(s1, s2, s3)
		}
	}
	return false
}

// 判定是否是回文
func isOk(s string) bool {
	if len(s) <= 1 {
		return true
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
