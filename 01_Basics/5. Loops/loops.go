package main

import "fmt"

func main() {
	/*Go lang only have for loops but they can work like diff loops*/

	//Traditional for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//While loop
	j := 0
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	//Range Based loop
	for k := range 5 {
		fmt.Println("range", k)
	}

	//Infinite Loop
	for {
		fmt.Println("Infinite Loop")
		break
	}

}
