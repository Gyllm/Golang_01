package main

import "fmt"

func main() {
	a := 1
	for a <= 100 {
		if a%3 == 0 {
			fmt.Println("fizz")
			if a%5 == 0 {
				fmt.Println("fizz buzz")
			}
		} else if a%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(a)
		}
		a = a + 1
	}
}
