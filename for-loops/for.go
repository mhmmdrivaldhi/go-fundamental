package main

import (
	"fmt"
)

func main() {
	// for loop
	var i = 0

	for i <= 10 {
		fmt.Println("Looping ke - ", i)
		i++
	}

	for index := 11; index <= 20; index++ {
		fmt.Println("Index ke - ", index)
	}

	// for range with slice
	var number = []int{1, 2 , 3, 4, 5}

	for i := 0; i < len(number); i++ {
		fmt.Println(number[i])
	}

	var names = []string{"Val", "Rap", "Bot", "Jot", "Yuk"}

	// for range using index
	for i, value := range names {
		fmt.Println("Index", i, " =", value)
	}

	// for range without index
	for _, value := range names {
		fmt.Println("Value =", value)
	}

	// for range with map
	var book = make(map[string]string)
	book["title"] = "5cm"
	book["author"] = "Donny Dhirgantoro"
	book["release_year"] = "2005"
	book["pages"] = "591 Pages"

	for key, value := range book {
		fmt.Println("Key =", key, "|", "Value =", value)
	}
}