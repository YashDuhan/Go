package main

import "fmt"

const name string = "Yash"

func main() {
	fmt.Println(name)

	//It will give error because name is in const right now, we can only create a new variable "name" to shadow it
	// name = "Ram"

	var name = "Raj" //It can shadow the name variable
	fmt.Println(name)

	/*
		//But it can't be declared again
		name := "Abhay"
		fmt.Println(name)
	*/
	name = "Abhay" //It can be edited but we can't create a new var with "name"
	fmt.Println(name)
}
