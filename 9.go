package main

import "fmt"

func main() {
	var balance, amount float64
	var action int
	fmt.Scan(&balance, &action)

	switch action {
	case 1:
		fmt.Println(balance)
	case 2:
		fmt.Scan(&amount)
		balance += amount
		fmt.Println(balance)
	case 3:
		fmt.Scan(&amount)
		if amount <= balance {
			balance -= amount
			fmt.Println(balance)
		} else {
			fmt.Println("Недостаточно средств")
		}
	}
}
