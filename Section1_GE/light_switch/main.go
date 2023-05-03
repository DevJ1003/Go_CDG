package main

import (
	"fmt"
)

var lightSwitchIsOn bool = false

func main() {

	printlightSwitchState()
	togglelightSwitch()
	printlightSwitchState()
	togglelightSwitch()
	printlightSwitchState()

}

// ===================================================================
func printlightSwitchState() {

	fmt.Println("The light switch is off: ", lightSwitchIsOn)

}

// ===================================================================
func togglelightSwitch() {

	lightSwitchIsOn = !lightSwitchIsOn

}
