package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Start")
	// fmt.Println("Proccess")
	// defer fmt.Println("Done")

	// var num = 4

	// defer fmt.Println("Result Defer :", num)

	// num += 10
	// num *= 2

	// fmt.Println("Result Num :", num)

	// var num1 = 5
	// var num2 = 0
	
	// defer fmt.Println("Done")
	// fmt.Println("Start")
	// fmt.Println(num1 / num2)

	// defer add(8, multiply(5, 3))
	// add(3, 4)

	// fmt.Println("Start")
	// defer fmt.Println("Done")
	// defer add(2, 5)
	// defer add(5, 5)
	// defer multiply(3, 2)

	fmt.Println("Start")
	defer loop()
	fmt.Println("Done")

}

func add(num1 int, num2 int) {
	result := num1 + num2
	fmt.Println(result)
}

func multiply (num1 int, num2 int) int {
	result := num1 * num2
	fmt.Println(result)
	return result
}

func loop() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done Looping")
}
