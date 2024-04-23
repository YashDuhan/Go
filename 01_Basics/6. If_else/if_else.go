package main

import "fmt"

func main() {
	i := 7
	if i%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

	j := 13
	if j%2 == 0 || j%3 == 0 {
		fmt.Println(j, "is divisible by both 2 and 3")
	} else if i%2 == 0 {
		fmt.Println(j, "is only divisible by 2")
	} else if i%3 == 0 {
		fmt.Println(j, "is only divisible by 3")
	} else {
		fmt.Println(j, "is neither divisible by 2 nor by 3")
	}

}
