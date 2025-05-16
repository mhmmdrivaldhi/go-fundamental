package main

import (
	"fmt"
)

func main() {
	var numbers = []int{4, 3, 7, 11}
	var anotherNumbers = numbers

	fmt.Println("Numbers :", numbers)
	fmt.Println("Another Numbers :", anotherNumbers)

	numbers[1] = 9

	fmt.Println("Numbers :", numbers)
	fmt.Println("Another Numbers :", anotherNumbers)

	fmt.Println("=============================")

	var scores = []int {7, 8, 6, 9}

	multiplyBy10(scores)
	fmt.Println("Value Scores :", scores)
}

func multiplyBy10(numbers []int){
	for i := range numbers {
		numbers[i] = numbers[i] * 10
	}

	fmt.Println("Numbers di func :", numbers)
}