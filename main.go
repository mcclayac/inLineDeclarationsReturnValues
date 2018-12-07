package main

import (
	"fmt"
	"os"
)

func main () {


	//Option #1
	if numberBytes, theError := fmt.Printf("Hello World Option 1!\n"); theError != nil {
		os.Exit(1)
	} else {
		fmt.Printf("Printed %d bytes\n\n", numberBytes)
	}
	//Number of bytes has the scope of the if statements

	// Option #2
	numberBytes := 0
	var theError error

	if numberBytes, theError = fmt.Printf("Hello World Option 2!\n"); theError != nil {
		os.Exit(1)
	}

	fmt.Printf("Printed %d bytes\n\n", numberBytes)

	// Number of bytes has the scope of the if statements

	// Option #3
	numberBytes2, theError2 := fmt.Printf("Hello World  Option3 !\n")
	if theError2 != nil {
		os.Exit(1)
	} else if numberBytes2 > 100 {
		fmt.Printf("Printed %d bytes\n\n", numberBytes2)
	}

	// Option 4
	_, theError3 := fmt.Printf("Hello World Option 4!\n")
	if theError3 != nil {
		os.Exit(1)
	}

	// Option 5
	if _, theError4 := fmt.Printf("Hello World Option #5"); theError4 != nil {
		os.Exit(1)
	}




}

