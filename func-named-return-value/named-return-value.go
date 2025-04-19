package main

import "fmt"

func main() {
	fullname, age, married := introMySelf()

	fmt.Printf("Hello my name is %s, My age is %d, And my status is %t (Not Married)", fullname, age, married)
}

func introMySelf() (fullname string, age int, married bool) {
	fullname = "Muhammad Rivaldhi"
	age = 20
	married = false

	return
}