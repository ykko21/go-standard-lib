package main

import (
	"fmt"
	"strings"
)

func main1() {
	s := "We live in Closter in New Jersey."
	fmt.Println(strings.Contains(s, "Closter"))
	fmt.Println(strings.ContainsAny(s, "esz"))
	fmt.Println(strings.ContainsAny(s, "qz"))
	fmt.Println(strings.Index(s, "in"))
	s = "zzkqg"
	vowel := "aeiouAEIOU"
	fmt.Println(strings.IndexAny(s, vowel))
	s = "We live in Closter in New Jersey."
	fmt.Println(strings.HasPrefix(s, "We"))
	fmt.Println(strings.HasSuffix(s, "Jersey."))
	fmt.Println(strings.Count(s, "in"))
}
