package main

import (
	"flag"
	"fmt"
)

// Main function!
func main() {

	var gopherName string
	flag.StringVar(&gopherName, "gophername", "Gopher", "The name of the Gopher")
	flag.Parse()
	fmt.Println("Hello " + gopherName + "!")

}
