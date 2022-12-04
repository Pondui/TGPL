// Write a program wordfreq to report the frequency of each word in an input text file.
// Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into words instead of lines.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	test := wordfreq("testfile")
	fmt.Println(wordfreq("testfile"))
	fmt.Println(test[""])
}

// Returns map[string]int
func wordfreq(filename string) map[string]int {
	wordCount := make(map[string]int)
	content, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file: %s/n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(content)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		b := scanner.Bytes()
		runeSlc := []rune(string(b))

		// exclude punctuation and empty strings
		for i := 0; i < len(runeSlc); i++ {
			if unicode.IsPunct(runeSlc[i]) {
				copy(runeSlc[i:], runeSlc[i+1:])
				runeSlc = runeSlc[:len(runeSlc)-1]
				i--
			}
		}

		b = []byte(string(runeSlc))
		// Exclude empty strings
		if string(b) != "" {
			wordCount[strings.ToLower(string(b))]++
		}
	}

	return wordCount
}
