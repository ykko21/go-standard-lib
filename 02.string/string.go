package main

import (
	"fmt"
	"strings"
)

func main1() {
	t1()
}

func t1() {
	str := "Hello World"
	for i, ch := range str {
		fmt.Println(i, string(ch))
	}
}

func t2() {
	fmt.Println("dog" > "cat")
	fmt.Println("dog" > "hen")
	fmt.Println(strings.Compare("dog", "cat"))
	fmt.Println(strings.Compare("dog", "hen"))
	fmt.Println(strings.Compare("dog", "dog"))
	fmt.Println(strings.EqualFold("Hello", "hello"))
	s := "tHis is a good eXamPle."
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.Title(s))
	fmt.Println(strings.Title(strings.ToLower(s)))
	fmt.Println(strings.ToTitle(s))
}
