package greetingspackage

import (
	"fmt"
)

// indicating to Go that we want to export a function by upper casting function's first letter!
func PrintGreetings() {
	fmt.Println("I'm printing a message from PrintGreetings() function!")
}

// this function is unexported
// since it's unexported it will only be visible to that are within greetingspackage
func printGreetingsUnexported() {
	fmt.Println("I'm printing a message from the printGreetingsUnexported() function!")
}
