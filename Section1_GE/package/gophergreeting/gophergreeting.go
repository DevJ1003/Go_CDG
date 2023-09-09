package greetingspackage

import "fmt"

var MagicNumber int

// This is an exported function and can be called outside the package
func GopherGreetings() {

	fmt.Println("A very hollo to gophers! I'm printing from GopherGreetings() function!")
	// Now we are calling the unexported package
	// because this function is within the same package, wa have access to call it

}

// Example of Packages init() function
func init() {
	MagicNumber = 108
}
