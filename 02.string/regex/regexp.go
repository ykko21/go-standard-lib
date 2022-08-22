package main

import (
	"fmt"
	"log"
	"regexp"
)

// func main() {
func main() {
	fmt.Println("Hello")
	regex, err := regexp.Compile("^[a-zA-Z0-9]+$")
	if err != nil {
		log.Fatal("error")
	}
	fmt.Println(regex.MatchString("ab3jakeAb9"))
	fmt.Println(regex.MatchString("ab3jake!Ab9"))
}
