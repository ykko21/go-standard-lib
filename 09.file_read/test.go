package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//t3_()
	t3()
}

func t3_() {
	tf, err1 := os.Create("hello_copy.txt")
	if err1 != nil {
		if err1 == os.ErrExist {
			log.Fatal("File already exists")
		} else {
			log.Fatal(err1)
		}
	}
	var offset int64 = 0
	t1 := []byte("Hello ")
	t2 := []byte("World")
	n, err := tf.WriteAt(t1, offset)
	handleError(err)
	offset += int64(n)
	n, err = tf.WriteAt(t2, offset)
	tf.Sync()
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func t3() { //copy antoher way using io.Copy
	f, err := os.Open("hello.txt")
	if err != nil {
		if err == os.ErrNotExist {
			log.Fatal("File not exist")
		} else {
			log.Fatal(err)
		}
	}
	const bufferSize = 20
	defer f.Close()

	tf, err1 := os.Create("hello_copy.txt")
	if err1 != nil {
		if err1 == os.ErrExist {
			log.Fatal("File already exists")
		} else {
			log.Fatal(err1)
		}
	}

	buf := make([]byte, bufferSize)
	var offset int64 = 0
	for {
		//buf := make([]byte, bufferSize)
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("Reached to EOF!")
				fmt.Print(string(buf))
				//tf.Sync()
				return
			} else {
				log.Fatal(err)
			}
		}
		tf.WriteAt(buf[:n], offset)
		offset += int64(n)
	}
}

func t2() {
	const BuffSize = 20
	f, _ := os.Open("d:/workspaces/learning/go_standard_lib/09.file_read/hello.txt")
	defer f.Close()

	buf := make([]byte, BuffSize)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		fmt.Println("Byte read: ", n)
		fmt.Println("Content: ", string(buf[:n]))
	}
}

func t1() {
	cont, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cont))
}
