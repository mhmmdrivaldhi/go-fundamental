package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Rivaldhi"
	names[2] = "Bogor"

	fmt.Println("First Name : ", names[0])
	fmt.Println("Last Name : ", names[1])
	fmt.Println("Full Name : ", names[0] + " ", names[1])
	fmt.Println("Address : ", names[2])

	var students = [3]string{"Muhammad Rivaldhi", "Reza Abdullah", "Riska Apriani"}

	fmt.Println("Student 1 : ", students[0])
	fmt.Println("Student 2 : ", students[1])
	fmt.Println("Student 3 : ", students[2])
	fmt.Println("All Student : ", students)

	// array func
	fmt.Println("==========================")
	fmt.Println(" Array Func ")
	fmt.Println(len(students))
	fmt.Println(len(students[0]))
}