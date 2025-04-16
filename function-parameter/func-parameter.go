package main

import "fmt"

func main() {
	sayHelloTo("Muhammad Rivaldhi", 20)
}

func sayHelloTo(name string, age int) {
	fmt.Printf("Hello My name's = %s, and my age's = %d", name, age)
}