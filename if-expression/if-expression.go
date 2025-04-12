package main

import "fmt"

func main() {
	var username = "mhmmdrivaldhi"
	var password = "rivaldhi10"

	// if-else expression

	if username == "mhmmdrivaldhi" && password == "rivaldhi10" {
		fmt.Println("Login Successfully!")
	} else	{
		fmt.Println("Username or Password is Wrong!")
	}

	fmt.Println("=========================================")

	// if-else-elseif expression and using short statement

	var number = 1


	if i := 10; i <= number {
		fmt.Println(" i Smaller Than the number is false")
	} else if i == number {
		fmt.Println("i Bigger Than the number is true")
	} else {
		fmt.Println("Overall Is Wrong!")
	}

}