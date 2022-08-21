package main

import (
	"fmt"
	"reflect"
	"unicode"
)

func main() {

	s := "ab1dCDEf!Ged123"

	for i, c := range s {
		v := string(rune(c))
		fmt.Println(
			i, c,
			reflect.TypeOf(c), v,
			unicode.IsDigit(c), unicode.IsLower(c),
		)
		//v := rune(c)
		//fmt.Println(v)
	}

	fmt.Println("")
}
