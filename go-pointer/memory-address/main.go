package main

import "fmt"

func main() {
	var names = "Rivaldhi"

	fmt.Println("Name :", names)
	fmt.Println("Alamat Memory dari name :", &names)

	var age = 20
	fmt.Println("Age :", age)
	fmt.Println("Alamat Memory dari age :", &age)
}

