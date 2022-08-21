package main

import (
	"fmt"
	"log"
)

func main() {
	//a, b := t1()
	//fmt.Println(a, b)
	t2()
}

type person struct {
	FirstName string
}

func (p *person) printFirstName() {
	fmt.Println(p.FirstName)
}

func t2() {
	var p1 person
	p1.FirstName = "Peter"

	p2 := person{
		FirstName: "James",
	}

	log.Println(p1.FirstName)
	log.Println(p2.FirstName)

	p1.printFirstName()
	p2.printFirstName()
}

func t1() (string, string) {
	return "Hello", "World"
}
