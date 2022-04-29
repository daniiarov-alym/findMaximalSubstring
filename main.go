package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const helpMessage = "The program finds " +
	"the length of maximal substring of input string " +
	"that contains at most 2 distinct characters.\n" +
	"Examples:\n" +
	"ababab -> 6 (ababab)\n" +
	"abbca -> 3 (abb)\n" +
	"abaccab -> 4 (acca)"

	
//The function returns the length of maximal substring of input containing at most limit distinct symbols
func findMaximalSubstring(input string, limit uint) (size uint) {
	if limit == 0 {
		return 0
	}
	table := make(map[rune]uint)
	runes := []rune(input)
	start, end := uint(0), uint(0)

	for ; end < uint(len(runes)); end++ {
		table[runes[end]]++
		if uint(len(table)) > limit {
			for uint(len(table)) > limit {
				if table[runes[start]] == 1 {
					delete(table, runes[start])
				} else {
					table[runes[start]]--
				}
				start++
			}
		}
		if end-start >= size {
			size = end - start + 1
		}
	}
	return size
}

func main() {
	fmt.Println(helpMessage)
	scanner := bufio.NewScanner(os.Stdin)
	var maxDistinct uint = 2
	for scanner.Scan() {
		input := scanner.Text()
		if scanner.Err() != nil {
			_, _ = fmt.Fprintf(os.Stderr, "unexpected error while reading input: %s\n", scanner.Err())
			continue
		}
		if len(input) == 0 {
			continue
		}
		words := strings.Fields(input)
		if len(words) == 0 {
			continue
		}
		for _, word := range words {
			maxLength := findMaximalSubstring(word, maxDistinct)
			fmt.Printf("The length of maximal substring of a string %s with at most %d distinct characters is %d\n",
				word,
				maxDistinct,
				maxLength)
		}
	}
}
