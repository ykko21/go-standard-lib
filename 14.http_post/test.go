package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	t2()
}

func t2() {

	data := url.Values{}
	data.Add("zipcode", "12345")
	data.Add("state", "New Jersey")

	const url = "https://httpbin.org/post"
	resp, err := http.PostForm(url, data)
	if err != nil {
		log.Fatal(err)
	}
	content, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(string(content))
}

func t1() {
	const url = "https://httpbin.org/post"
	reqBody := strings.NewReader(`
	{
		"first_name":"James",
		"last_name":"Koon"
		"email":"james.koon@gmail.com
	}
	`)
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		log.Fatal(err)
	}
	content, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(string(content))
}
