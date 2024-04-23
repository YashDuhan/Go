package main

import "fmt"

func add(a int, b int) { // No return type means void
	result := a + b
	printRes(result)
}

func printRes(result int) {
	fmt.Println("Result is :", result)
}

func plus(a int, b int) int { // return type int
	return a + b
}

func check(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}

}

func main() {
	add(3, 4)
	fmt.Println(plus(10, 20))
	if check(5) {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is Odd")
	}
}
