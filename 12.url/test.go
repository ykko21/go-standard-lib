package main

import (
	"fmt"
	"net/url"
)

func main() {
	t2()
}

func t2() {
	newurl := &url.URL{
		Scheme:   "https",
		Host:     "www.example.com",
		Path:     "/account",
		RawQuery: "name=peter&zipcode=12345",
	}
	fmt.Println(newurl.String())

	param := url.Values{}
	param.Add("x", "3")
	param.Add("y", "5")
	newurl.RawQuery = param.Encode()
	fmt.Println(newurl.String())
}

func t1() {
	s := "https://www.example.com:8000/user?name=james"

	r, _ := url.Parse(s)
	fmt.Println(r.Scheme)
	fmt.Println(r.Host)
	fmt.Println(r.Path)
	fmt.Println(r.Port())
	fmt.Println(r.RawQuery)

	qry := r.Query()
	fmt.Println(qry["name"])
}
