// Rewrite reverse to use an array pointer instead of a slice.

package main

import "fmt"

func main() {
	fmt.Println(reverse(&[10]int{0,1,2,3,4,5,6,7,8,9}))
}

// An array's length is part of its type and must be specified in the function params
func reverse(arr *[10]int) [10]int {
    for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		// Automatic dereferencing of arr when [] is used
        arr[i], arr[j] = arr[j], arr[i]
    }
	return *arr
}