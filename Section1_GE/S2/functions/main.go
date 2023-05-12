package main

import (
	"fmt"
)

var y = 4

// =========================================================
func OddOrEven(x int) {

	if x%2 == 0 {
		fmt.Println("The number,", x, ", is even.")
	} else {
		fmt.Println("The number,", x, ", is odd.")
	}
	if y%2 == 0 {
		fmt.Println("The number,", y, ", is even.")
	} else {
		fmt.Println("The number,", x, ", is odd.")
	}

}

// =========================================================
func SumOfTwoNumbers(x int, y int) int {
	return x + y
}

// =========================================================
func SumAndDiffOperationsOnTwoNumbers(x, y int) (sum int, diff int) {
	return x + y, x - y
}

// =========================================================
func MultiSums(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}

	return sum
}

// =========================================================
func main() {

	var NumberToCheck = 7
	OddOrEven(NumberToCheck)
	fmt.Println()

	fmt.Println("The sum of 10 and 8:", SumOfTwoNumbers(10, 8))
	fmt.Println()

	sum, difference := SumAndDiffOperationsOnTwoNumbers(10, 8)
	fmt.Printf("Input: 10 and 8, The sum: %d, The diff: %d\n", sum, difference)
	fmt.Println()

	fmt.Println("Multi Sum result:", MultiSums(0, 5, 7, 6, 2, 3, 4, 8, 1, 9))
	fmt.Println()

	func() {
		fmt.Println("Hi, I'm an anonymous function. They call me that, because I don't have a name!")
	}()

}