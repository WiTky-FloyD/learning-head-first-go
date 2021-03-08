package main

import "fmt"

func main() {
	fmt.Printf("The %s cost %d cents each. \n", "gumballs", 23)
	fmt.Printf("That will be $%f please.\n", 0.23 * 5)

	/*
		%f Floating point number
		%d Decimal integer
		%s String
		%t Boolean
		%v Any value (chooses an appropriate format based on the supplied value’s type)
		%#v Any value, formatted as it would appear in Go program code
		%T Type of value
		%% A literal percent sign
	*/

	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An interger: %d\n", 15)
	fmt.Printf("A string: %s\n", "hello")
	fmt.Printf("A boolean: %t\n", false)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n")

	fmt.Printf("%v %v %v", "", "\t", "\n")
	fmt.Printf("%#v %#v %#v", "", "\t", "\n")
}