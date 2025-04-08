package main

import "fmt"

func main() {
	const first_name string = "Muhammad"
	const last_name string = "Rivaldhi"
	const age int = 20
	const is_marriage bool = false

	fmt.Println("First Name = ", first_name)
	fmt.Println("Last Name = ", last_name)
	fmt.Println("Is Marriage = ", is_marriage)
	fmt.Println("Age = ", age)

	fmt.Println("===============================")
	fmt.Println("Constant with multiple variable")

	const (
		country string = "Indonesia"
		graduation_year int = 2023
		school_graduate string = "SMK Plus Pelita Nusantara"
		vocational_field string = "Software Engineering"
	)

	fmt.Println("Contry = ", country)
	fmt.Println("Graduation Year = ", graduation_year)
	fmt.Println("School Gradute = ", school_graduate)
	fmt.Println("Vocational Field = ", vocational_field)
}