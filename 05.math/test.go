package main

import (
	"fmt"
	"math"
)

func main() {
	t3()
}

func t3() {

}

func t2() {
	x := 10.0
	fmt.Println(math.Abs(-x))
	fmt.Println(math.Pow(x, 2))
	fmt.Println(math.E, math.Exp(2))
}

func t1() {
	fmt.Println(math.Pi)
	fmt.Println(math.Ceil(math.Pi))
	fmt.Println(math.Floor(math.Pi))
	fmt.Println(math.Trunc(math.Pi))
	a := float64(5)
	b := float64(3)
	fmt.Println(math.Min(a, b))
	fmt.Println(math.Max(a, b))

	fmt.Println(int(a) % int(b))
	fmt.Println(math.Mod(a, b))

	fmt.Println(math.Round(11.5))
	fmt.Println(math.Round(-11.5))
	fmt.Println(math.RoundToEven(11.5))
	fmt.Println(math.RoundToEven(-11.5))
}
