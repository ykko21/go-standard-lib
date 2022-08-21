package main

import (
	"fmt"
	"strings"
)

func main1() {
	transform := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			v := int(r) + 2
			if v > int('Z') {
				v = int('A') + (v % int('Z')) - 1
			}
			return rune(v)
		case r >= 'a' && r <= 'z':
			v := int(r) + 2
			if v > int('z') {
				v = int('a') + (v % int('z')) - 1
			}
			return rune(v)
		}
		return r
	}
	s := "abz123ABYZ"
	newStr := strings.Map(transform, s)
	fmt.Println(newStr)
}
