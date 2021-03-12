package main

import "fmt"

func main() {
	myInt := 4
	myIntPointer := &myInt
	fmt.Println(*myIntPointer)

	myFloat := 98.6
	myFloatPointer := &myFloat
	fmt.Println(*myFloatPointer)

	myBool := true
	myBoolPointer := &myBool
	fmt.Println(*myBoolPointer)
	line()

	myInt2 := 4
	fmt.Println(myInt2)
	myInt2Pointer := &myInt2
	*myInt2Pointer = 8
	fmt.Println(*myInt2Pointer)
	fmt.Println(myInt2)
}

func line()  {
	fmt.Println("==========================")
}