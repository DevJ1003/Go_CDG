package main

import (
	"fmt"
	socialmedia "std/Projects/Go_CDG/Section1_GE/social_mp/socialmedia"
)

// Main function!
func main() {

	myPost := socialmedia.NewPost("DevJoshi", socialmedia.Moods["thrilled"], "Go is awesome!", "Check out the Go website!",
		"https://golang.org", "", "", []string{"go", "golang", "programming language"}, []string{})
	fmt.Printf("myPost: %+v\n", myPost)

}
