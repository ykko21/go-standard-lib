package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	t5()
}

func t5() { //shuffle
	rand.Seed(time.Now().UnixNano())
	const numstring = "one two three four five"
	words := strings.Fields(numstring)
	fmt.Println(words)
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)
}

func t4() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		c := string('A' + rune(rand.Intn(5)))
		fmt.Println(c)
	}
}

func t3() {
	rand.Seed(time.Now().UnixNano())
	s := []string{"apple", "banana", "grapes", "kiwi", "orange"}
	indexes := rand.Perm(len(s))
	for i := 0; i < len(indexes); i++ {
		fmt.Println(s[indexes[i]])
	}
}

func t2() {
	rand.Seed(time.Now().UnixNano())
	s := []string{"apple", "banana", "grapes", "kiwi", "orange"}
	fmt.Println(s[rand.Intn(5)])
}

func t1() {
	rand.Seed(time.Now().UnixNano())
	//rand.Seed(1234)
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(3))

	fmt.Println(rand.Float32())
	fmt.Println(rand.Float64())
}
