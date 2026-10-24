package main

import "fmt"

func main() {
	var a, b int
	var operation string

	if _, err := fmt.Scan(&a, &b); err != nil {
		return
	}

	if _, err := fmt.Scan(&operation); err != nil {
		return
	}

	switch operation {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль")
		} else {
			fmt.Println(a / b)
		}
	default:
		fmt.Println("Ошибка: неизвестная операция.")
	}
}
