package main

import "fmt"

func main() {
	var a int

	if _, err := fmt.Scan(&a); err != nil {
		return
	}

	if a >= 0 {
		fmt.Println("Положительное")
	} else {
		fmt.Println("Отрицательное")
	}

	if a%2 == 0 {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}

	if a%3 == 0 && a%5 == 0 {
		fmt.Println("Делится и на 3, и на 5")
	} else {
		fmt.Println("Не делится одновременно на 3 и 5")
	}

}
