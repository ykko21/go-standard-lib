package main

import "fmt"

type circle struct {
	radius int
	border int
}

func main() {
	//t1()
	//t2()
	t3()
}

func t3() {
	f := 123.4567
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%10f\n", f)
}

func t2() {
	c := circle{
		radius: 20,
		border: 5,
	}
	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)
	fmt.Printf("%T\n", c)
}

func t1() {
	x := 20
	f := 123.45873641
	fmt.Printf("%d\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%t\n", x > 10)

	fmt.Printf("%f\n", f)
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%e\n", f)
}
