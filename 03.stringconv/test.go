package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func main() {
	t2()
}

func t2() { //parse
	b, _ := strconv.ParseBool("TRUE") //True, TRUE, true, ...
	fmt.Println(b)
	f, _ := strconv.ParseFloat("12.34", 16)
	fmt.Println(f)

	r1 := strconv.FormatBool(true)
	fmt.Println(r1)

	r2 := strconv.FormatInt(23, 10)
	fmt.Println(r2)

	r3 := strconv.FormatFloat(12.23, 'E', -1, 32)
	fmt.Println(r3)

	r4 := strconv.FormatFloat(12.23, 'f', -1, 32)
	fmt.Println(r4)
}

func t1() {
	n := 100
	fmt.Println(string(rune(n)))
	fmt.Println(strconv.Itoa(n), reflect.TypeOf(strconv.Itoa(n)))
	fmt.Println(strconv.Itoa(n), reflect.TypeOf(strconv.Itoa(n)))

	//v, err := strconv.Atoi("20a0")
	v, err := strconv.Atoi("200")
	if err != nil {
		log.Fatal("error")
	}
	fmt.Println(v)
	fmt.Println(reflect.TypeOf(v))
}
