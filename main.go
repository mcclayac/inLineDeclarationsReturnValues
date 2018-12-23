package main

import (
	"fmt"
	"os"
)

func main() {

	//Option #1
	if numberBytes, theError := fmt.Printf("Hello World #1!\n"); theError != nil {
		os.Exit(1)
	} else {
		fmt.Printf("Printed num %d of bytes\n\n", numberBytes)
	}
	//Number of bytes has the scope of the if statements

	// Option #2
	numberBytes2 := 0
	var theError2 error

	if numberBytes2, theError2 = fmt.Printf("Hello World Option 2!\n"); theError2 != nil {
		os.Exit(1)
	}

	fmt.Printf("Printed %d bytes  #2\n\n", numberBytes2)

	// Number of bytes has the scope of the if statements

	// Option #3
	numberBytes3, theError3 := fmt.Printf("Hello World  #3 !\n")
	if theError3 != nil {
		os.Exit(1)
	} else if numberBytes3 > 100 {
		fmt.Printf("Printed %d bytes\n\n", numberBytes3)
	}

	// Option 4
	_, theError4 := fmt.Printf("Hello World Option #4!\n")
	if theError4 != nil {
		os.Exit(1)
	}

	// Option 5
	if _, theError5 := fmt.Printf("Hello World Option #5"); theError5 != nil {
		os.Exit(1)
	}
}
