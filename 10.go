package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 0 {
		return
	}

	var sum, min, max, even, odd int

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)

		sum += x

		if i == 0 || x > max {
			max = x
		}
		if i == 0 || x < min {
			min = x
		}

		if x%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	average := float64(sum) / float64(n)
	fmt.Println(sum, average, max, min, even, odd)
}
