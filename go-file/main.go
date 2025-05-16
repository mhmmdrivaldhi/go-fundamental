package main

import (
	"fmt"
	"strconv"
)

var names []string

func main() {
	var input string
	fmt.Println("Names :", names)
	fmt.Print("Masukkan Jumlah Nama : ")
	fmt.Scanln(&input)

	len, _ := strconv.Atoi(input)
	for i := 0; i < len; i++ {
		fmt.Print("Enter The Names :")
		fmt.Scanln(&input)

		names = append(names, input)
	}
	fmt.Println("Names :", names)
}