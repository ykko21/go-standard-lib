package main

import (
	"fmt"
)

func main() {
	t3()
}

func t1() {
	m := make(map[string]string)
	m["first_name"] = "Peter"
	m["last_name"] = "Papa"
	fmt.Println(m)

	//m = make(map[string]interface)
}

type User struct {
	FirstName string
	LastName  string
}

func t2() {
	m := make(map[string]User)

	u1 := User{
		FirstName: "Peter",
		LastName:  "Pompa",
	}

	m["user"] = u1

	fmt.Println(m)
}

func t3() {
	m := make(map[string]interface{})
	m["name"] = "Peter"
	m["age"] = 20
	m["married"] = false
	fmt.Println(m)
}
