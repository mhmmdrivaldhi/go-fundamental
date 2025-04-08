package main

import "fmt"

func main() {
	var val32 int32 = 100000
	var val64 int64 = int64(val32)
	var val8 int8 = int8(val32) // akan menjadi min (-), karena tidak bisa menerima nilai dari int32

	fmt.Println(val32)
	fmt.Println(val64)
	fmt.Println(val8)

	fmt.Println("===================")

	var name = "Rivaldhi"
	var e byte = name[0] // e adalah tipe data byte
	var eString = string(e) // eString akan meng konversi tipe data dari byte menjadi string lagi

	fmt.Println(e)
	fmt.Println(eString)
}