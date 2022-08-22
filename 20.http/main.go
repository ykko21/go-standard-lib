package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, World!")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Request URI: ", r.RequestURI)
		fmt.Println("Bytes written: ", n)
	})
	http.HandleFunc("/welcome", greeting)
	port := "8081"
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Listening at %s ...\n", port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func greeting(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Welcome to ykko!!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Request URI: ", r.RequestURI)
	fmt.Println("Bytes written: ", n)
}
