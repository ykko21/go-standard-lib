package main

import (
	"fmt"
)

func main() {
	t1()
}

func t1() {
	f := 123.4567
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%10f\n", f)
	fmt.Printf("%10.2f\n", f)
	fmt.Printf("%+10.2f\n", f)
	fmt.Printf("%010.2f\n", f)
}

func t2() {
	f := 10.23412
	v := fmt.Sprintf("%.2f", f)
	fmt.Println(v)

	v = fmt.Sprintf("[%10f]\n", f)
	fmt.Println(v)
}
