package main

import "fmt"

func main()  {
	// x := 4 initailzatime statement
	// x <= 6 condition expression
	// x++ post statement
	for x := 4; x <= 6; x++ {
		fmt.Println("x is now", x)
	}
	line()

	for t := 3; t > 0; t-- {
		fmt.Println(t)
	}
	line()

	for x := 1; x <= 5; x += 2 {
		fmt.Println(x)
	}
	line()

	for x := 15; x >= 5; x -= 5 {
		fmt.Println(x)
	}
	line()

	x := 1 
	for x <= 3 {
		fmt.Println(x)
		x++
	}
	line()

	y := 3 
	for y >= 1{
		fmt.Println(y)
		y--
	}
}

func line() {
	fmt.Println("=============================")
}