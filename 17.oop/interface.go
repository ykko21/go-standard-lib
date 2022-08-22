package main

import "fmt"

type Animal interface {
	Say() string
}

type Dog struct {
	Name string
}

func (d *Dog) Say() string {
	return "Woof"
}

type Gorilla struct {
	Name string
}

func (g *Gorilla) Say() string {
	return "krrr"
}

func main() {
	d := Dog{
		Name: "Serong",
	}
	g := Gorilla{
		Name: "Kingkong",
	}
	fmt.Println(d.Name, d.Say())
	fmt.Println(g.Name, g.Say())
}
