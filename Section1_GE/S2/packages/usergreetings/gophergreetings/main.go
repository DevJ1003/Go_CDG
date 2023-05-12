package greetingspackage

import (
	"fmt"
	// "Projects/Go_CDG/Section1_GE/packages/usergreetings/greetings"
)

var MagicNumber int

// ===========================================================
func GopherGreetings() {

	fmt.Println("A very hello my fellow gophers! I'm printing from the GopherGreetings() function!")
	// printGreetingsUnexported()

}

// func printGreetingsUnexported() {
// 	fmt.Println("hjahahahahjaha")
// }

// ===========================================================
func init() {

	MagicNumber = 108

}
