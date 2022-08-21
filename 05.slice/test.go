package main

import (
	"fmt"
	"sort"
)

func main() {
	t1()
}

func t1() {
	var s1 []int
	s1 = append(s1, 1)
	s1 = append(s1, 3)
	s1 = append(s1, 2)

	fmt.Println(s1)
	sort.Ints(s1)
	fmt.Println(s1)
	sort.Sort(sort.Reverse(sort.IntSlice(s1)))
	fmt.Println(s1)

	s2 := []string{"a", "b", "c"}
	sort.Sort(sort.Reverse(sort.StringSlice(s2)))
	fmt.Println(s2)
}
