package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"flag"
	"fmt"
)

// The ordering matters, you must specify the flags before the argument
// this works: 		./ex4.2 -sha384 banana
// but not this:    ./ex4.2 banana -sha384

func main() {
	sha384Ptr := flag.Bool("sha384", false, "returns sha384 encoding")
	sha512Ptr := flag.Bool("sha512", false, "returns sha512 encoding")
	
	input := flag.Arg(0)
	flag.Parse()

	if *sha384Ptr {
		byteArr := sha512.Sum384([]byte(input))
		encodePrint(byteArr[:])
	}
	if *sha512Ptr {
		byteArr := sha512.Sum512([]byte(input))
		encodePrint(byteArr[:])
	}
	if !*sha384Ptr && !*sha512Ptr {
		byteArr := sha256.Sum256([]byte(input))
		encodePrint(byteArr[:])
	}
}

func encodePrint(byteSlice []byte) {
	str := base64.StdEncoding.EncodeToString(byteSlice)
	fmt.Println(str)
}
