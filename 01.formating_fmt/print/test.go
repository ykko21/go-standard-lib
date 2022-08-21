package main

import "fmt"

func main() {
	//var_const()
	//slice()
	//string_test()
	//int_test()
	slice2()
}

func slice2() {
	items := []int{1, 2, 3, 4}
	length, err := fmt.Println(items)
	fmt.Println(length, err)
}

func string_test() {
	s := "hello"
	fmt.Printf("value:%v\t\t, memory address:%p\n", s, &s)
	s = "hello" + " world"
	fmt.Printf("value:%v\t, memory address:%p\n", s, &s)
}

func int_test() {
	i1 := 1
	fmt.Printf("value:%v, memory address:%p\n", i1, &i1)
	i1 = i1 + 2
	fmt.Printf("value:%v, memory address:%p\n", i1, &i1)
}

func slice1() {
	items := []int{1, 2, 3, 4}
	fmt.Printf("items:%v, memory address:%p\n", items, &items)
	items = append(items, 5)
	fmt.Printf("items:%v, memory address:%p\n", items, &items)
}

func var_const() {
	var a int
	var b, c float32
	if true {
		a = 1
		b, c = 1.1, 2.2
	} else {
		a = 2
		b, c = 2.1, 3.2
	}
	fmt.Println(a, b, c)
	a = 3
	b, c = 4.1, 5.6
	fmt.Println(a, b, c)

	const i int = 10
	const j, k string = "jj", "kk"
	print(i, ",", j, ",", k)
	//i = 20 //compile error
	//j, k = "jjj", "kkk" //compile error
}
