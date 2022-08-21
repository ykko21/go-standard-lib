package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	t3()
}

func t3() { //http client with setting timeout
	//const url = "https://httpbin.org:5000/get"
	const url = "https://httpbin.org/get"

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	buf := make([]byte, resp.ContentLength)
	byteCount, err := resp.Body.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(byteCount, string(buf))
}

func t2() {
	const url = "https://httpbin.org/get"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("Status", resp.Status)
	fmt.Println("Status Code", resp.StatusCode)
	fmt.Println("Protocol", resp.Proto)
	fmt.Println("Content Length", resp.ContentLength)
	var sb strings.Builder
	content, _ := ioutil.ReadAll(resp.Body)
	bytecount, _ := sb.Write(content)
	fmt.Println(bytecount, sb.String())
}

func t1() {
	const url = "https://httpbin.org/get"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("Status", resp.Status)
	fmt.Println("Status Code", resp.StatusCode)
	fmt.Println("Protocol", resp.Proto)
	fmt.Println("Content Length", resp.ContentLength)
	buf := make([]byte, resp.ContentLength)
	resp.Body.Read(buf)
	fmt.Println(string(buf))
}
