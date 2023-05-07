package main

import (
	"flag"
	"fmt"
)

func main() {

	var gopherName string

	flag.StringVar(&gopherName, "gophername", "Gopher", "The name of the Gopher")
	flag.Parse()
	fmt.Println("Hello " + gopherName + "!")

}

/*

1st input ==> go run main.go ................. output ==> Hello Gopher!

2nd input ==> go run main.go -help

3rd input ==> go run main.go -gophername Dev ................. output ==> Hello Dev!

*/
