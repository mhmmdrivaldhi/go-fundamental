package main

import "fmt"

func main() {
	var name1 = "Rivaldhi"
	var name2 = "Rivaldhi"
	var full_name = "Muhammad Rivaldhi"

	var result bool = name1 == name2
	fmt.Println("Result 1 : ", result)

	var result2 = name1 == full_name
	fmt.Println("Result 2 : ", result2)

	var result3 = name2 != full_name
	fmt.Println("Result 3 : ", result3)

	var a = 10
	var b = 20
	var c = a < b
	fmt.Println("Result 4 : ", c)

	var res = a > b
	fmt.Println("Result 5 : ", res)
}