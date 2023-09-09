package main

import (
	"fmt"
)

// Main function!
func main() {

	for i := 0; i <= 10; i++ {

		// if(i==0){

		// 	continue
		// }
		fmt.Println(i)

	}

	var x int = 0

	if x%2 == 0 {
		fmt.Println("\nX is even!")
	} else if x%2 != 0 {
		fmt.Println("\nX is odd!")
	} else {
		fmt.Println("\nX is 0!")
	}

}
