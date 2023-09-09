package main

import (
	"fmt"
	greetingspackage "std/Projects/Go_CDG/Section1_GE/package/greetingspackage"
)

func main() {

	greetingspackage.PrintGreetings()
	greetingspackage.GopherGreetings()

	fmt.Println("The value of the magic number : ", greetingspackage.MagicNumber)

}
