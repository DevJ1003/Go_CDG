package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

const UsernameRegex string = `^@?(\w){1,15}$`

func main() {

	var UsernameInput string
	flag.StringVar(&UsernameInput, "username", "Gopher", "The GopherFace Username To Check")
	flag.Parse()

	fmt.Println("GopherFace Username Validation Checker")
	fmt.Println("Checking Syntax for username, \"", UsernameInput, "\", resulted in: ", CheckUsernameSyntax(UsernameInput), "\n")

}

func CheckUsernameSyntax(username string) bool {

	validationResult := false
	r, err := regexp.Compile(UsernameRegex)
	if err != nil {
		log.Fatal(err)
	}

	validationResult = r.MatchString(username)
	return validationResult
}

/*

1st input ==> go run main.go -username ^Java_Duke^ ............... output ==> bool false

2nd input ==> go run main.go -username @EngineerDev ............... output ==> bool true

*/
