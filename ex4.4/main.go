// Write a version of rotate that operates in a single pass
// The question doesn't say in constant space

package main

import "fmt"

func main() {
	fmt.Println(rotate([]int{1,2,3,4,5}, 3))
	fmt.Println(rotate([]int{1,2,3,4,5}, -3))
	fmt.Println(rotate([]int{1,2,3,4,5}, -9))
}

// positive n signifies right rotaion, negative n signifies left rotation
func rotate(s []int, n int) []int {
	result := make([]int, len(s))

	for i, el := range s {
		j := newIndex(i, n, len(s))
		result[j] = el
	}
	return result
}

func newIndex(i int, n int, length int) int {
	j := i + n
	if j >= length {
		j = j % length
	} else if j <= -length {
		j = j % length
	}

	if j < 0 {
		j = j + length
	}
	return j
}
