package main

import "fmt"

type machine struct {
	tipe string
	cc int
}

type transport struct {
	brand string
	year  int
	model string
	price int
	machine
}

func main() {

	var a = transport{
		brand: "Honda",
		year: 2021,
		model: "Mobilio",
		price: 2000000000,
		machine: machine{
			tipe: "Type R",
			cc: 2999,
		},
	}
	fmt.Println("Embeded Cars :", a)
	fmt.Println("Machine-Type :", a.machine.tipe)
	/*
	var a transport
	a.brand = "Mclaren"
	a.year = 2019
	a.model = "Senna"
	a.price = 14000000000

	fmt.Println("Cars :", a)
	fmt.Println("Brand :", a.brand)
	fmt.Println("Years :", a.year)
	fmt.Println("Type :", a.model)
	fmt.Println("Price :", a.price)

	fmt.Println("========================")

	var b = transport{}
	b.brand = "Alphard"
	b.year = 2020
	b.model = "Veelfire"
	b.price = 7000000000

	fmt.Println("Cars 2 :", b)

	fmt.Println("========================")

	var c = transport{"Toyota", 2018, "Avanza Veloz", 65000000}
	fmt.Println("Cars 3 :", c)

	fmt.Println("========================")

	var d = transport{model: "Brio Hytec", brand: "Honda", year: 2017, price: 10500000}

	fmt.Println("Cars 4 :", d)

	fmt.Println("-- Struct BY Value --")

	var e = transport{brand: "Honda", year: 2018, model: "Jazz", price: 1200000000}
	var f transport = e

	fmt.Println("Before Change Value")
	fmt.Println("Cars e :", e)
	fmt.Println("Cars y :", f)
	fmt.Printf("Address e : %p\n", &e)
	fmt.Printf("Address f : %p\n", &f)
	
	fmt.Println("After Change Value")
	f.model = "Brio"
	fmt.Println("Cars e", e)
	fmt.Println("Cars f", f)

	// fmt.Println("=====================")

	// var nt = transport{brand: "Toyota", model: "Supra", year: 2022, price: 50000000000}
	// updateTransport(nt)
	// fmt.Println("Transport in function main :", nt)

	// fmt.Println("-- Pass BY Refrence --")

	// var cars *transport = &nt

	// fmt.Printf("Address nt : %p\n", &nt)
	// fmt.Printf("Address cars : %p\n", &cars)

	// cars.model = "Civic"
	// fmt.Println("Refrence nt :", nt)
	// fmt.Println("Refrence Cars :", cars)

	fmt.Println("============================")

	var z = transport{brand: "Honda", year: 2022, model: "Mobilio", price: 20000000}

	updateTransport(&z)
	fmt.Println("Transport in func main :", &z)

	fmt.Println("=============================")
	var k = new(transport)
	fmt.Printf("Address K : %p\n", k)
	*/
}

func updateTransport(newTransport *transport){
	newTransport.brand = "Toyota"
	newTransport.model = "Sigra"
	newTransport.price = 35000000
	newTransport.year = 2021

	fmt.Println("Update Transport :", newTransport)
}