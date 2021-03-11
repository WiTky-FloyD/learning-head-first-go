package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("This is an error")
	fmt.Println(err.Error())
	// log.Fatal(err)

	err2 := fmt.Errorf("This is an error")
	fmt.Println(err2.Error())
	fmt.Println(err2)
}
