package main

import (
	"fmt"
)

func main() {

	var myInteger int32
	fmt.Println("Value of myInteger: ", myInteger)

	var myFloatingPointNumber float64 = 36.333
	var myOtherFloatingPointNumber float64 = 306.96969696
	fmt.Println("Sum of my floating point numbers: ", myFloatingPointNumber+myOtherFloatingPointNumber)

	x, y, z := 0, 1, 2
	fmt.Printf("x: %d\ty: %d\tz: %d\n", x, y, z)

	myComplexNumber := 5 + 24i
	fmt.Println("Value of myComplex Number: ", myComplexNumber)

	// const (
	// 	speedOfLight      = 299792458
	// 	pi                = 3.14
	// 	myFavouriteNumber = 108
	// )

	var (
		a int = 0
		b     = 1.8 + 3i
		c     = 2.7
	)

	fmt.Printf("a: %v\tb: %v\tc: %v\n", a, b, c)

}
