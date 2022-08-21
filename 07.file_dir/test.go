package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	t2()
}

func t2() {
	stat, err := os.Stat("D:/workspaces/learning/go_standard_lib/07.os/test.go")
	if err != nil {
		log.Fatal("no file exists")
	}
	fmt.Println(stat.Mode())
	fmt.Println(stat.IsDir())
	fmt.Println(stat.ModTime())
	fmt.Println(stat.Size())
}

func t1_(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func t1() {
	println(t1_("./test.go"))
	println(t1_("./test1.go"))
	println(t1_("D:/workspaces/learning/go_standard_lib/07.os/test.go"))
}
