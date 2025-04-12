package main

import "fmt"

func main() {
	student := map[string]string{
		"name":    "Muhammad Rivaldhi",
		"address": "Bogor",
	}

	// Add Data Map
	student["work"] = "Backend Developer"

	fmt.Println("Data Map of Student : ", student)
	fmt.Println("Data Student by Map key : ", student["name"])
	fmt.Println("Data Student by Map key : ", student["address"])
	fmt.Println("Data Map of student After Add data", student)
	fmt.Println("Data Lenght in Map : ", len(student))

	delete(student, "address")
	fmt.Println("Data student After Delete : ", student)
	fmt.Println("Data Lenght Student after Delete : ", len(student))

	fmt.Println("================================")
	
	var book = map[string]string{
		"title": "Madilog",
		"author": "Tan Malaka",
		"release_year": "2000",
		"pages": "441 Pages",
	}

	fmt.Println("Details of Book : ")
	fmt.Println("Book Title : ", book["title"])
	fmt.Println("Book Author : ", book["author"])
	fmt.Println("Book Release Year : ", book["release_year"])
	fmt.Println("Book Pages : ", book["pages"])

	book["price"] = "59000"
	fmt.Println("Book Price : ", book["price"])
}