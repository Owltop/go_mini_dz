package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 0 {
		return
	}

	var x float64
	fmt.Scan(&x)
	min, max := x, x

	for i := 1; i < n; i++ {
		fmt.Scan(&x)

		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}

	fmt.Println(max, min)
}
