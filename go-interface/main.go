package main

import "fmt"

type Shape interface {
	getArea() float64
	getPerimeter() float64
}

type Rectangle struct {
	width  float64
	lenght float64
}

func (r Rectangle) getArea() float64 {
	return r.width * r.lenght
}

func (r Rectangle) getPerimeter() float64 {
	return 2 * (r.width + r.lenght)
}

type Square struct {
	side float64
}

func (s Square) getArea() float64 {
	return s.side * s.side
}

func (s Square) getPerimeter() float64 {
	return 4 * s.side
}

// func getAreaOfRectangle(r Rectangle) {
// 	fmt.Println("Luas :", r.getArea())
// }

// func getAreaOfSquare(s Square) {
// 	fmt.Println("Luas :", s.getArea())
// }

func getArea(s Shape) {
	fmt.Println("Luas :", s.getArea())
}

func main() {
	var shape1 Shape = Square{side: 5}
	var shape2 Shape = Rectangle{width: 10, lenght: 15}
	var shape3 Shape = Rectangle{width: 7, lenght: 5}

	fmt.Printf("Shape 1 %#v\n", shape1)
	fmt.Printf("Shape 2 %#v\n", shape2)
	fmt.Printf("Shape 3 %#v\n", shape3)

	fmt.Println("=================================")

	shapes := []Shape{shape1, shape2, shape3}

	for _, shape := range shapes {
		fmt.Printf("%#v memiliki luas %.1f dan keliling %.1f\n", shape, shape.getArea(), shape.getPerimeter())
	}

	fmt.Println("==================================")

	// getAreaOfRectangle(Rectangle{width: 6, lenght: 6})
	// getAreaOfSquare(Square{side: 8})

	getArea(Rectangle{width: 3, lenght: 10})
	getArea(Square{side: 10})

	fmt.Println("==================================")

	var square1 Shape = Square{side: 6}

	fmt.Println("getArea :", square1.getArea())
	fmt.Println("get Perimeter :", square1.getPerimeter())

	var originalSquare Square = square1.(Square)
	fmt.Println("get Area :", originalSquare.getArea())
	fmt.Println("get Perimeter :", originalSquare.getPerimeter())
	fmt.Println("Side :", originalSquare.side)

	fmt.Println("=====================================")

	var anything any

	anything = "Lesmana"
	fmt.Printf("Anything  %T: %v\n", anything, anything)
	
	anything = 20
	fmt.Printf("Anything  %T: %v\n", anything, anything)
	
	anything = true
	fmt.Printf("Anything  %T: %v\n", anything, anything)
	
	anything = 37.5
	fmt.Printf("Anything  %T: %v\n", anything, anything)
	
	anything = []string{"Golang", "Laravel", "NextJS"}
	fmt.Printf("Anything  %T: %v\n", anything, anything)
	
}