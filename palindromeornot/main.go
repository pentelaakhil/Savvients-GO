package main

import "fmt"

func main() {
	var number, remainder, NUM int
	var reverse int = 0

	fmt.Print("Enter any integer : ")
	fmt.Scan(&number)

	NUM = number
	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10

		if number == 0 {
			break
			// Break Statement it comes out of the loop
		}
	}

	if NUM == reverse {
		fmt.Printf("%d is a Palindrome \n", NUM)
	} else {
		fmt.Printf("%d is not a Palindrome \n", NUM)
	}

}
