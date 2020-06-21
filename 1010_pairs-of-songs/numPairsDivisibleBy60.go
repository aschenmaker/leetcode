package main

import "fmt"

func main() {
	// time := []int{30, 20, 150, 100, 40, 60}
	time := []int{60, 60, 60}
	fmt.Println(numPairsDivisibleBy60(time))
}

func numPairsDivisibleBy60(time []int) int {
	counter := 0

	maps := make(map[int]int, 60)
	for _, v := range time {
		x := v % 60
		var sub int
		if x == 0 {
			sub = 0
		} else {
			sub = 60 - x
		}
		if _, ok := maps[sub]; ok {
			counter += maps[sub]
		}
		maps[x]++
	}
	fmt.Println(maps)
	return counter
}
