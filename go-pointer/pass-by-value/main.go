package main

import (
	"fmt"
)

func main() {
	var x = 4
	var y = x

	y = y + 1
	fmt.Println("x :", x)
	fmt.Println("y :", y)

	fmt.Println("Alamat Memory dari x :", &x)
	fmt.Println("Alamat Memory dari y :", &y)

	fmt.Println("=============================")

	var a = 3
	increase(a)

	fmt.Println("a :", a)
}

func increase(n int){
	n = n + 1

	fmt.Println("n :", n)
}