package main

import "fmt"

func main() {
	//basic decleration
	var a = "Hello"
	var b = "World"
	fmt.Println(a + b)

	//integer
	var c = 1
	var d = 2
	fmt.Println(c + d)

	//int
	var e, f int = 1, 2
	fmt.Println(e * f)

	// Will give error
	// var g = 3, h = 3

	var g int //value initialized with 0
	fmt.Println(g)

	// var h //Will return error as data type isn't mentioned

	//Directly creating variables
	h := "Hello Everyone"
	fmt.Println(h)

	i := 34
	j := 35
	fmt.Println(i + j)

	//creating 2 variables of diff datatypes
	var k, l string = "Mango", "Apples"
	fmt.Println("I have an", l, "and a", k)
	// var m,n int = 3, string = "Apple" will give error
	m, n := 4, "Apple"
	fmt.Println("He had", m, n)

	/*
		//concatenating diff data types will result in error
		o := 8
		p := "Wires"
		fmt.Println(o + p)
	*/
}
