package main

import "fmt"

func romanToInt(s string) int {
	tmp := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	output := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && tmp[s[i]] < tmp[s[i+1]] {
			output -= tmp[s[i]]
		} else {
			output += tmp[s[i]]
		}
	}
	return output
}
func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))

}
