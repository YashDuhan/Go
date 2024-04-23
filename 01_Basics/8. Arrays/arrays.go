package main

import "fmt"

func main() {
	//array
	var a [5]int

	//printing an empty array
	fmt.Println("Empty array : ", a)

	//adding element at an index
	a[0] = 10
	fmt.Println("Array after an insertion : ", a)

	//printing an index
	fmt.Println("Element at 0th index is : ", a[0])

	//length of an array
	length := len(a)
	fmt.Println("Length is ", len(a), "or", length)

	// Initialising an array with elements at the time of creation
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array b is : ", b)

	// 2-D Array
	twoDArr := [4][5]int{}
	// twoDArr := [4][5]int;
	/*
		twoDArr := [4][5]int; will not work
		either use
			var twoDArr[4][5]int or
			twoDArr := [4][5]int{}  Mandatory to initialise values {} represents 0
	*/
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			twoDArr[i][j] = i + j
		}
	}
	fmt.Println("2D array : ", twoDArr)
}
