package main

import (
	"fmt"
)

func main() {

	z := 36

	// switch statement
	switch z {
	case 1:
		fmt.Println("z is equal to 1")
	case 2:
		fmt.Println("z is equal to 2")
	case 3:
		fmt.Println("z is equal to 3")
	default:
		fmt.Println("z is not equal to 1, 2 or 3")
	}

}
