package main

import (
	"fmt"
)

var y = 4

func add(x int, y int) int {

	return x + y

}

func main() {

	var number = 4

	fmt.Println("The sum of x and y :", add(number, y))
}
