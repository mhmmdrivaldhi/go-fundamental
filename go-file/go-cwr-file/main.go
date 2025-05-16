package main

import (
	"bufio"
	"fmt"
	"os"
)

var path = "/Users/muham/enigmacamp/assets/"
var fileName = "data.txt"
var filePath = path + fileName

func main() {
	// var input string
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Print("Enter the names : ")
	// scanner.Scan()
	// input = scanner.Text()

	// writeFile(input)
	readFile()
}

func createFile() {
	_, error := os.Stat(filePath)
	if os.IsNotExist(error) {
		file, error := os.Create(filePath)
		if error != nil {
			fmt.Println(error)
			return
		}
		defer file.Close()

		fmt.Println("File", fileName, "Success to Created.")
	}
}

func writeFile(input string) {
	file, error := os.OpenFile(filePath, os.O_APPEND | os.O_RDWR, 0666)
	if error != nil {
		fmt.Println(error)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(input + "\n")
	writer.Flush()
}

func readFile() {
	file, error := os.OpenFile(filePath, os.O_APPEND | os.O_RDWR, 0666)
	if error != nil {
		fmt.Println(error)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

