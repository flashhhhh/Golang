package main

import "fmt"

type Rectangle struct {
	width  int
	length int
}

func (rec Rectangle) Area() int {
	return rec.width * rec.length
}

func (rec Rectangle) setWidth(width int) {
	rec.width = width
}

func (rec Rectangle) setLength(length int) {
	rec.length = length
}

func (rec *Rectangle) setWidthPointer(width int) {
	rec.width = width
}

func (rec *Rectangle) setLengthPointer(length int) {
	rec.length = length
}

func main() {
	rect1 := Rectangle{10, 5}
	rect2 := Rectangle{width: 10, length: 5}
	rect3 := Rectangle{length: 5}

	fmt.Println("Area of rect1 is: ", rect1.Area())
	fmt.Println("Area of rect2 is: ", rect2.Area())
	fmt.Println("Area of rect3 is: ", rect3.Area())

	// Set width and length doesn't work
	rect1.setWidth(20)
	rect1.setLength(20)
	fmt.Println("Area of rect1 is: ", rect1.Area())

	// Set width and length using pointer
	rect1.setWidthPointer(20)
	rect1.setLengthPointer(20)
	fmt.Println("Area of rect1 is: ", rect1.Area())
}