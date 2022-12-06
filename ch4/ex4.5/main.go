// Write an in-place function to eliminate adjacent duplicates in a []string slice

package main

import "fmt"

func main() {
	fmt.Println(dedupAdjacent([]string{"a", "a", "b", "c", "c", "d", "e", "e"}))
	fmt.Println(dedupAdjacent([]string{
		"banana",
		"banana",
		"banana",
		"apple",
		"apple",
		"orange",
		"pear",
		"strawberry",
		"kiwi",
		"kiwi",
		"pineapple",
		"mango",
		"mango",
	}))
}

func dedupAdjacent(s []string) []string {
	for i := 0; i < len(s)-1; {
		if s[i] == s[i+1] {
			copy(s[i+1:], s[i+2:])
			s = s[:len(s)-1]
		} else {
			i++
		}
	}

	return s
}
