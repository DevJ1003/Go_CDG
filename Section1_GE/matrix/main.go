package main

import (
	"fmt"
)

// Main function!
func main() {

	myMatrix := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

	fmt.Println("Value at [0][0]: ", myMatrix[0][0])
	fmt.Println("Value at [1][1]: ", myMatrix[1][1])
	fmt.Println("Value at [2][2]: ", myMatrix[2][2])

	fmt.Printf("%+v\n", myMatrix)
	fmt.Println("\n")

	printThreeByFourMatrix(myMatrix)

}

func printThreeByFourMatrix(inputMatrix [3][4]int) {

	rowLength := len(inputMatrix)
	columnLength := len(inputMatrix)

	for i := 0; i < rowLength; i++ {
		for j := 0; j < columnLength; j++ {
			fmt.Printf("%5d ", inputMatrix[i][j])
		}
		fmt.Println()
	}

}
