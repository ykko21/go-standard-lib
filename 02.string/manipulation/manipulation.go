package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	t6()
}

func t6() {
	a := "ac.!@#b$5%^&e*(4)k1kkz?"
	f := func(c rune) bool {
		r, err := regexp.MatchString("[a-zA-Z0-9]", string(c))
		if err != nil {
			log.Fatalf("error1: %v", err)
		}
		return !r
	}
	result := strings.FieldsFunc(a, f)
	fmt.Printf("%q", result)
}

func t5() {
	a := "ac.!@#b$5%^&e*(4)k1kkz?"
	regex, err := regexp.Compile("[a-zA-Z0-9]")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	result := regex.FindAllString(a, -1)
	fmt.Printf("%q", result)
}

func t4() {
	a := "ac.!@#$%^&*()?"
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], a[i] == 'a', unicode.IsPunct(rune(a[i])), " ")
		fmt.Printf("%c\n", a[i])
	}
}

func t3() {
	s := "He said \"It's a beautiful day!\""
	arr := strings.Fields(s)
	fmt.Println(arr)

	f := func(c rune) bool {
		return unicode.IsPunct(c)
	}

	arr = strings.FieldsFunc(s, f)
	fmt.Printf("%q", arr)
}

func t2() {
	arr := []string{"one", "two", "three"}
	s := strings.Join(arr, " - ")
	fmt.Println(s)
}

func t1() {
	fmt.Println("Hello")
	s := "He said \"It's a beautiful day!\""
	s1 := strings.Split(s, " ")
	fmt.Println(s1)
	fmt.Printf("%v\n", s1)
	fmt.Printf("%+v\n", s1)
	fmt.Printf("%q\n", s1)
}
