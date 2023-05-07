package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var repeatCount int
	var err error

	if len(os.Args) >= 2 {
		repeatCount, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < repeatCount; i++ {
			fmt.Println("Hello Gopher!")
		}
	}
}

/*

1st input ==> go run main.go 9

2nd input ==> go run main.go 108

3rd input ==> go run main.go 108 | wc -l

*/
