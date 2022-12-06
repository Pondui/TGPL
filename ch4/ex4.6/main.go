// Write an in-place function that squashes each run of adjacent Unicode spaces
// (see unicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	str := "abc  def  hij \tklm \nopq \n"
	// Prints the string without interpreting escapes, expect "abc def hij klm opq "
	s := fmt.Sprintf("%#v", string(squashUnicodeSpace([]byte(str))))
	fmt.Println(s)
}

func squashUnicodeSpace(b []byte) []byte {
	size := 0

	// First dedup unicode spaces
	for size < len(b)  {
		r1, size1 := utf8.DecodeRune(b[size:])
		r2, size2 := utf8.DecodeRune(b[size+size1:])
		if unicode.IsSpace(r1) && unicode.IsSpace(r2) {
			copy(b[size:], b[size+size2:])
			b = b[:len(b)-size2]
			// Replace non space unicode.IsSpace characters (i.e. tabs, newlines) with spaces
			if newR1, newSize1 := utf8.DecodeRune(b[size:]); newR1 != ' ' {
				b[size] = []byte(" ")[0] // Works only because a space is 
				copy(b[size+1:], b[size+newSize1:])
				b = b[:len(b)-newSize1+1]
			}
			// Prints the string without interpreting escapes
			// s := fmt.Sprintf("%#v", string(b))
			// fmt.Println(s)
		} else {
			size += size1
		}
	}
	return b
}