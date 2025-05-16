package main

import "fmt"

func main() {
	//Anonymous Function
	func() {
		fmt.Println("Hello World!")
	}()

	// Anonymous Function in Variable
	halo := func ()  {
		fmt.Println("Halo Dunia")
	}

	halo()

	// Passing Argument in Anonymous Function
	func (word string)  {
		fmt.Println(word)	
	} ("Hello, Rivaldhi")

	// Passing Argument in Variable function
	hello := func(name string) {
		fmt.Println("Hello", name)
	} 
	hello("Rivaldhi")

	// Passing Anonymous function as Argument
	greetEnglish := func(name string) string {
		return "HI, " + name
	}

	greetIndonesian := func (name string) string {
		return "Hai, " + name
	}

	greetItalian := func (name string) string  {
		return "Ciao, " + name
	}

	greet("Marissa", greetEnglish)
	greet("Thompson", greetIndonesian)
	greet("Marco", greetItalian)

	add := func (num1 int, num2 int) int {
		return num1 + num2
	}

	multiply := func (num1 int, num2 int) int {
		return num1 * num2
	}

	calculate(3,2, add)
	calculate(6,7, multiply)
}

func greet(name string, f func(name string)string){
	fmt.Println(f(name))
} 

func calculate(a int, b int, operator func(x int, y int) int ) {
	fmt.Println(operator(a, b))
}