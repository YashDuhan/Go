/*The range form of the for loop iterates over a slice*/
package main

func main() {
	arr := []int{10, 20, 30, 40, 50, 60, 70}
	/*
		A range function returns 2 values
			i) Index of the element
			ii) Copy of that element
	*/
	for i, v := range arr {
		println("Element present at", i, "index is", v)
	}
}
