package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	t1()
}

func t1() {
	s := "abc"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i], reflect.TypeOf(s[i]), string(s[i]))
	}

	base := int('a')
	fmt.Println(base)

	var sb strings.Builder
	fmt.Println("Capacity: ", sb.Cap())
	for i := 0; i < 26; i++ {
		sb.WriteString(string(rune(base + i)))
	}
	fmt.Println("Capacity: ", sb.Cap())
	fmt.Println(sb.String())
	sb.Grow(1024)
	fmt.Println("Capacity: ", sb.Cap())
	base = int('A')
	for i := 0; i < 26; i++ {
		sb.WriteString(string(rune(base + i)))
	}
	fmt.Println("Capacity: ", sb.Cap())

	for i := 0; i <= 10; i++ {
		sb.WriteString(fmt.Sprintf("String %d ", i))
	}
	fmt.Println(sb.String())
	sb.Reset()
	fmt.Println(sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
}
