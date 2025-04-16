package main

import "fmt"

func main() {
	name, age := introSelf("Muhammad Rivaldhi", 19)
	fmt.Println(name, "Your age's ", age)
}

func introSelf(name string, age int) (string, int) {
	if name == "" || age != 20 {
		return "Hello Sir", 0
	} else {
		return "Hello! " + name, age
	}
}