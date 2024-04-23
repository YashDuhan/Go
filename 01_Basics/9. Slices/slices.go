package main

import "fmt"

func main() {
	/*
		Slices are similar to vectors in C++
		Their size can be shrinked and decreased based on need
	*/
	/*
		Syntax: slice_name := []datatype{values}
	*/

	// Creating an empty slice with length 0 and capacity 0
	// In an empty array the vacant elements are represented via [0,0,0] but in slices it's represented as []
	myslice := []int{}
	fmt.Println("Array :", myslice, "\nSize of Array :", len(myslice))

	//creating a slice with elements in it
	my_slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Array :", my_slice, "\nSize of Array :", len(my_slice))

	// usage of cap() function
	// cap is used to find the number of elements an array can have
	arr := [3]int{1}
	fmt.Println("Array :", arr, "\nSize of Array :", len(arr), "\nCapacity of Array :", cap(arr))

	// capacity of previous slices
	fmt.Println("Capacity of myslice :", cap(myslice))
	fmt.Println("Capacity of my_slice :", cap(my_slice))

	// string slice
	string_slice := []string{"A", "Go", "Slice", "Can", "Have", "Many", "Elements"}
	fmt.Println(string_slice)

	//Best Method to create a slice
	/*
		It's always recommended to use a make() function while creating a slice; a make() function is like a new function of C++
		It's used to allocate memory
	*/
	s := make([]int, 5, 8) // here 5 is the size and 8 is the capacity
	fmt.Println("Size of slice s :", len(s), "\nCapacity of slice s :", cap(s), "\nSlice s :", s)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	s[3] = 4
	s[4] = 5
	// s[5] = 6 (Index out of range error)
	fmt.Println("Size of slice s :", len(s), "\nCapacity of slice s :", cap(s), "\nSlice s :", s)

	//Appending elements
	a := []int{}
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	fmt.Println(a)

	//Appending elements in a slice created via make function
	b := make([]int, 3)
	fmt.Println("Size :", len(b), "Capacity :", cap(b))
	b = append(b, 1)
	fmt.Println("Size :", len(b), "Capacity :", cap(b))
	fmt.Println(b) //appending an element can increase the capacity by double and size by 1
}
