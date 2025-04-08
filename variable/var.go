package main

import "fmt"

func main() {
	var name string

	name = "Muhammad"
	fmt.Println(name)

	name = "Rivaldhi"
	fmt.Println(name)

	var full_name = "Muhammad Rivaldhi"
	fmt.Println(full_name)

	var age = 20
	fmt.Println(age)

	fmt.Println("=====================")
	fmt.Println("Declare Variable Using :=")

	school_graduate := "SMK Plus Pelita Nusantara"
	fmt.Println(school_graduate)

	graduation_year := 2023
	fmt.Println(graduation_year)

	fmt.Println("======================")
	fmt.Println("Declare with multiple variable")

	var (
		first_name = "Muhammad"
		last_name = "Rivaldhi"
		full_Name = "Muhammad Rivaldhi"
		country = "Indonesia"
		birth_date = 10072004
	)
	fmt.Printf("First Name = %s, Last Name = %s, Full Name = %s, Country = %s, Birthdate = %d", first_name, last_name, full_Name, country, birth_date)

}