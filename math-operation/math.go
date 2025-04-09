package main

import "fmt"

func main() {
	// math operation basic
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a + b
	fmt.Println("Result : ", c)

	fmt.Println("===============")

	// augmented assignment
	var i = 10
	 i *= 10
	fmt.Println("Augmented : ", i)

	fmt.Println("===============")

	// Unary Opertor
	i++
	fmt.Println("Uanry Operator : ", i)

	var negative = -20
	var positif = 10
	res := negative + positif
	fmt.Println("Res : ", res)
}