package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var b1 bool
	fmt.Println("Bool zero value", b1)
	b2 := true
	fmt.Println(reflect.TypeOf(b2))

	var i1 int
	fmt.Println("Integer ero value", i1)
	i2 := 3
	fmt.Println(reflect.TypeOf(i2))
	
	var f1 float64
	fmt.Println("Float ero value", f1)
	f2 := 1.0
	fmt.Println(reflect.TypeOf(f2))

	var s1 string
	fmt.Println("String zero value", s1)
	s2 := "Hello, Types"
	fmt.Println(reflect.TypeOf(s2))
}