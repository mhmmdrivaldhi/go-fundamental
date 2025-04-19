package main

import "fmt"

func main() {
	amount := multiply(5, 10, 15, 20)
	fmt.Println(amount)
}

func multiply(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total *= number
	}
	return total
}