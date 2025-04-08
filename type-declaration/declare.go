package main

import "fmt"

func main() {
	type NoKTP string
	type marrige_status bool

	var nik_rivaldhi NoKTP = "123456789"
	fmt.Println(nik_rivaldhi)

	var status marrige_status = false
	fmt.Println(status)
}