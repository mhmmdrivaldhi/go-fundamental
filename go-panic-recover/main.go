package main

import "fmt"

func main() {
	var input string
	fmt.Print("Enter Your Name :")
	fmt.Scanln(&input)
	 
	fmt.Println("Start")
	if !isEmpty(input) {
		fmt.Println(input)
	}
	fmt.Println("Done")
}

func isEmpty(input string) (empty bool){
	defer func ()  {
		if r := recover(); r != nil {
			fmt.Println("Terjadi Panic : ", r)
			empty = true
		}
	}()

	if input == "" {
		panic("Input Is Must Required!")
	}
	empty = false
	return
}