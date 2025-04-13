package main

import "fmt"

func main() {
	// break statement
	for i := 0; i < 10; i++ {
		fmt.Println("I ke -", i)
		if i == 5 {
			break
		}
	}

	// continue statement
	for index := 0; index <= 10; index++ {
		if index % 2 == 0 {
			fmt.Println("Index ke -", index)
			continue
		}
	}
}