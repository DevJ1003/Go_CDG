package main

import (
	"fmt"
)

func main() {

	const (
		LosAngeles = 1984 + (iota * 4)
		Seoul
		Barcelona
		Atlanta
		Sydney
		Athens
		Beijing
		London
		Rio
		Tokyo
	)

	fmt.Println("These cities hosted or will host the Summer Olympics in the giver year...")
	fmt.Printf("%-18s %-18s \n", "City", "Year")
	fmt.Printf("%-18s %-18s \n", "LosAngeles", LosAngeles)
	fmt.Printf("%-18s %-18s \n", "Atlanta", Atlanta)
	fmt.Printf("%-18s %-18s \n", "Tokyo", Tokyo)

}
