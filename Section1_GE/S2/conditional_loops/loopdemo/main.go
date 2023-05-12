package main

import (
	"fmt"
	"time"
)

func main() {

	// classic for loop
	for i := 0; i <= 10; i++ {

		if i == 0 {
			continue
		}

		fmt.Println("Inside classic for loop, value of i: ", i)
	}

	fmt.Println("\n\n")

	// single condition for loop
	j := -20
	for j != 0 {
		fmt.Println("Inside single condition for loop, value of j: ", j)
		j++
	}

	fmt.Println("\n\n")

	loopTimer := time.NewTimer(time.Second * 5)
	// infinite loop, will break out in 5 seconds
	for {

		fmt.Println("Inside the infinite loop!")

		<-loopTimer.C
		break
	}

}
