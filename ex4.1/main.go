// Write a function that counts the number of bits that are different in two SHA256 hashes

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	str1 := "test string 1"
	str2 := "test string 2"
	sha1 := sha256.Sum256([]byte(str1))
	sha2 := sha256.Sum256([]byte(str2))
	fmt.Println(DiffBits(sha1, sha2))
}

func DiffBits(sha1, sha2 [32]byte) int {
	diffBits := 0
	for i := 0; i < len(sha1); i++ {
		shaXOR := sha1[i] ^ sha2[i]
		diffBits += PopCount(shaXOR)
	}
	return diffBits
}

func PopCount(x uint8) int {
	count := 0
	for x > 0 {
		// clears rightmost set bit
		x = x&(x-1)
		count ++
	}
	return count
}