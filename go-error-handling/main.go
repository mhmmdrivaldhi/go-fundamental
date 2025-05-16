package main

import (
	"errors"
	"fmt"
)

func main() {
	var result, err = divide(5, 5)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println("Result :", result)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("pembagian dengan angka 0")
	}
	return a / b, nil
}