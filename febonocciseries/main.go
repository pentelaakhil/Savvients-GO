package main

import "fmt"

func main() {
	i := 0
	j := 1
	k := 1
	for k <= 1000 {
		fmt.Print(k, ",")
		k = i + j
		i = j
		j = k
	}
	fmt.Println()

}
