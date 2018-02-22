package main

import "fmt"

func main() {
	if 72%2 == 0 {
		fmt.Println("72 is even")
	} else {
		fmt.Println("72 is odd")
	}

	if 84%4 == 0 {
		fmt.Println("84 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}