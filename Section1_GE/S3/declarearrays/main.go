package main

import (
	"fmt"
)

// 1st attempt at populating the array
func populateIntegerArray(input [5]int) {

	input[0] = 3
	input[1] = 6
	input[2] = 9
	input[3] = 12
	input[4] = 15

}

// 2nd attempt at populating the array with return value
func populateIntegerArrayWithReturnValue(input [5]int) [5]int {

	input[0] = 3
	input[1] = 6
	input[2] = 9
	input[3] = 12
	input[4] = 15
	return input

}

// =========================================================
func LinkinParkArrayExample() {

	// declare and initialize values in an array
	var LinkinPark [6]string
	LinkinPark[0] = "Chester"
	LinkinPark[1] = "Mike"
	LinkinPark[2] = "Joe"
	LinkinPark[3] = "Rob"
	LinkinPark[4] = "Brad"
	LinkinPark[5] = "Phoenix"

	fmt.Printf("LinkinPark consists: %v\n", LinkinPark)
	fmt.Println("The name of the first member in the LinkinPark: ", LinkinPark[0])

	fmt.Println("Length of LinkinPark: ", len(LinkinPark))

	var greatRockBandsFromThe90s [6]string
	greatRockBandsFromThe90s = LinkinPark
	fmt.Printf("Members from the great rock band from the 1990s: %v\n", greatRockBandsFromThe90s)

	fmt.Printf("LinkinPark member address: %p\n", &LinkinPark)
	fmt.Printf("greatRockBandsFromThe90s member address: %p\n", &greatRockBandsFromThe90s)
	if LinkinPark == greatRockBandsFromThe90s {
		fmt.Println("The LinkinPark array equals the greatRockBandsFromThe90s array\n\n")
	}

}

// =========================================================
func u2ArrayExample() {

	u2 := [4]string{"Bono", "Edge", "Adam", "Larry"}
	fmt.Printf("u2 consists: %v\n", u2)
	fmt.Println("The name of second band member in the u2 array: ", u2[1])

	fmt.Println("Length of u2: ", len(u2))

}

// =========================================================
func main() {

	var myArray [5]int
	fmt.Printf("Contents of myArray: %v\n\n", myArray)

	populateIntegerArray(myArray)
	fmt.Printf("Contents of myArray: %v\n\n", myArray)

	myArray = populateIntegerArrayWithReturnValue(myArray)
	fmt.Printf("Contents of myArray: %v\n\n", myArray)

	fmt.Println("Length of myArray: ", len(myArray))
	fmt.Println("===================================================")

	LinkinParkArrayExample()

	u2ArrayExample()

}
