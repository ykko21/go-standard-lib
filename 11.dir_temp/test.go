package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	t2()
}

func t2() { // temp dir
	tmpDir, err := ioutil.TempDir("", "tempdir_*")
	if err != nil {
		panic(err)
	}
	fmt.Println(tmpDir)
	defer os.RemoveAll(tmpDir)
}

func t1() { // temp file
	tempPath := os.TempDir()
	fmt.Println(tempPath)

	tempContents := []byte("Hello World")
	tmpFile, err := ioutil.TempFile(tempPath, "tempfile_*.txt")
	if err != nil {
		panic(err)
	}

	n, err := tmpFile.Write(tempContents)
	if err != nil {
		panic(err)
	}
	fmt.Println(tmpFile.Name())
	fmt.Println(n, "bytes written")

	data, _ := ioutil.ReadFile(tmpFile.Name())
	fmt.Println("Contents is ..")
	fmt.Println(string(data))
}
