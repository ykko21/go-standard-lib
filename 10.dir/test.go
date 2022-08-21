package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	t4()
}

func t4() {
	contents, _ := ioutil.ReadDir(".")
	for _, b := range contents {
		fmt.Println(b.Name(), b.IsDir())
	}
}

func t3() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	dir, _ = os.Executable()
	fmt.Println(dir)

}

func t2() {
	os.Remove("testdir01")
	os.RemoveAll("a/b/c")
	os.RemoveAll("a/b")
	os.RemoveAll("a")
}

func t1() { //createDir
	os.Mkdir("testdir01", 0755) //mode 4:read, 2:write, 1:exe 6(4+2)
	os.MkdirAll("a/b/c", os.ModePerm)
}
