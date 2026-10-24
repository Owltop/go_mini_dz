package main

import "fmt"

func printTable(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(i*j, " ")
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	printTable(n)
}
