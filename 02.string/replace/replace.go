package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "ababcdea"
	rep := strings.NewReplacer("a", "-", "b", "*")
	s1 := rep.Replace(s)
	fmt.Println(s1)
}
