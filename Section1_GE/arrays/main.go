package main

import (
	"fmt"
)

func populateIntegerArray(input [5]int) {

	input[0] = 3
	input[1] = 6
	input[2] = 9
	input[3] = 12
	input[4] = 15
}

func classexample() {

	var class [5]string
	class[0] = "John"
	class[1] = "Dave"
	class[2] = "Tim"
	class[3] = "Mike"
	class[4] = "Joe"

	fmt.Printf("The class consistes of %v students\n", class)
	fmt.Println("The name of the third student of the class is", class[2])
	fmt.Println("The length of class array is", len(class))

}

// Main function!
func main() {

	var myArray [5]int
	fmt.Printf("Contents of myArray: %v\n\n", myArray)

	populateIntegerArray(myArray)
	fmt.Printf("Contents of myArray: %v\n\n", myArray)

	classexample()

}
