package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	positive, negative, zeros := 0, 0, 0

	for i := 0; i < n; i++ {
		var x float64
		fmt.Scan(&x)

		if x > 0 {
			positive++
		} else if x < 0 {
			negative++
		} else {
			zeros++
		}
	}

	fmt.Println(positive, negative, zeros)
}
