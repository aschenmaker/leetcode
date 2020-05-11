package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickPow(x, n)
	}
	return 1 / quickPow(x, n)
}

func quickPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickPow(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func main() {
	var x float64 = 3.1
	n := 2
	ret := myPow(x, n)
	fmt.Println(ret)
}
