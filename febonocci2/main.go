package main

import "fmt"

func main() {
	i := 0
	j := 1
	k := 1

	var numb int
	fmt.Print("Fibonacci Numbers")
	fmt.Print("\n \n")
	fmt.Print("Given a number: \n")
	fmt.Scan(&numb)

	fmt.Print("Fibonacci Series :\n")

	for k <= numb {
		fmt.Print(k, ",")
		k = i + j
		i = j
		j = k
	}

}
