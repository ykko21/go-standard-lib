package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	t3()
}

func t3() { //append contents in an existing file
	f, err := os.OpenFile("hello.txt", os.O_APPEND, 0644)
	handleError(err)
	defer f.Close()
	n, err := f.WriteString("Good morning!\n")
	handleError(err)
	fmt.Println(n)
	f.Sync()
}

func t2() {
	f, err := os.Create("test1.txt")
	handleError(err)
	defer f.Close()

	fmt.Println(f.Name())
	f.WriteString("Hello!!!")

	data1 := []byte("Hello2!!!")
	f.Write(data1)

	f.Sync()
}

func t1() {
	data1 := []byte("Hello, World!\n")
	//ioutil.WriteFile("hello.txt", data1, 0644)
	ioutil.WriteFile("hello.txt", data1, 0644)

}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
