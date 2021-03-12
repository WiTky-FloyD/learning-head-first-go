package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))
	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool))
	line()

	var myInt2 int
	var myInt2Pointer *int
	myInt2Pointer = &myInt2
	fmt.Println(myInt2Pointer)

	var myFloat2 float64
	var myFloat2Pointer *float64
	myFloat2Pointer = &myFloat2
	fmt.Println(myFloat2Pointer)

	var myBool2 bool
	myBool2Pointer := &myBool2
	fmt.Println(myBool2Pointer)
}

func line()  {
	fmt.Println("===========================")
}
