package main

import "fmt"

func main() {
	var name = "Rivaldhi"

	switch name {
	case "Muhammad":
		fmt.Println("Not Exist")
	case "Rivaldi":
		fmt.Println("Tidak Ada")
	default:
		fmt.Println("Periksa kembali value yang anda masukkan!")
	}

	// switch expression with short statement

	var age = 20

	switch min_age := 17; min_age <= age {
	case true :
		fmt.Println("Umur memenuhi persyaratan")
	case false :
		fmt.Println("Umur Tidak Memenuhi persyaratan")
	}

	// switch nothing condition

	switch {
	case age > 20 :
		fmt.Println("Anda terlalu tua")
	case age < 20 :
		fmt.Println("Anda masih muda")
	default :
		fmt.Println("Umur anda sesuai")
	}
}