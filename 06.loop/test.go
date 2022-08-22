package main

import "fmt"

func main() {
	t1()
}

func t1() {

	type User struct {
		FirstName string
		Email     string
	}

	var users []User
	users = append(users, User{FirstName: "Peter", Email: "peter@gmail.com"})
	users = append(users, User{FirstName: "James", Email: "james@gmail.com"})
	users = append(users, User{FirstName: "Kal", Email: "kal@gmail.com"})

	for i, u := range users {
		fmt.Println(i, u)
	}
}
