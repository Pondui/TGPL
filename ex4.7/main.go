// Modify reverse to reverse the characters of a []byte slice that represents a UTF-8-encoded string, in place. Can you do it without allocating new memory?

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Test swapSubSlices
	str1 := "testString"
	str2 := "bananaPear"
	b1 := []byte(str1)
	b2 := []byte(str2)
	swapSubSlices(b1, 4)
	swapSubSlices(b2, 6)
	fmt.Println(string(b1))
	fmt.Println(string(b2))

	// Test reverse with non-ascii characters
	str3 := "你好 World 𝄞 ﭶ"
	b3 := []byte(str3)
	reverse(b3)
	fmt.Println(string(b3))
}

// Reverses []byte slice of a string, preserving utf8 encoding so that string is reversed
func reverse(b []byte) {
	if len([]rune(string(b))) <= 1 {
		return
	}

	// Calculate index at which to split
	_, lenI := utf8.DecodeRune(b)
	swapSubSlices(b, lenI)
	reverse(b[:len(b)-lenI])
	return
}

// Swaps two sub slices (i and j) within a slice, mutates the slice that is passed
// Left subslice has len i, right subslice has len j
// Will be used to account for various size utf8 encodings in []byte
func swapSubSlices(s []byte, lenI int) {
	lenJ := len(s) - lenI

	if len(s) <= 1 {
		return
	} else if len(s) == lenI {
		return
	}

	if lenI <= lenJ {
		// After loop, first lenI elements are correct
		for i, j := 0, lenI; i < lenI; i, j = i+1, j+1 {
			s[i], s[j] = s[j], s[i]
		}

		swapSubSlices(s[lenI:], lenI)
	} else {
		// After loop, last lenJ elements are correct
		for i, j := len(s)-1-lenJ ,len(s)-1; j >= lenI; i, j = i-1, j-1 {
			s[i], s[j] = s[j], s[i]
		}

		swapSubSlices(s[:lenI], lenI-lenJ)
	}
	return
}
