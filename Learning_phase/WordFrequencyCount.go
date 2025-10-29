package main

import (
	"fmt"
	"strings"
)


// Task 1: Word Frequency Count
// Write a function that takes a string of text and returns a map with the frequency count of each word in the text.
// The function should be case-insensitive and ignore punctuation.
func WordFrequencyCount(text string) map[string]int {
	words := strings.Split(strings.ToLower(text), " ")
	freq := make(map[string]int)
	for _, word := range words{
		if word != "" {
			freq[word]++
		}
	}
	return freq
}



// Task 2: Palindrome Checker
// Write a function that checks if a given string is a palindrome.

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	reversed := ""
	for i := len(s) - 1; i >= 0; i -- {
		reversed += string(s[i])
	}

	return s == reversed
}

func main() {

	// Task 1 test
	tests := []string {
		"Hello world hello",
		"Go is great. Go is fun!",
		"Test test TEST",
	}
	for _, test := range tests {
		fmt.Printf("Input: %q\n", test)
		fmt.Printf("Output: %v\n\n", WordFrequencyCount(test))
	}


	// Task 2 test
	palindromeTests := []string {
		"Racecar",
		"hello",
		"Madam",
		"Step on no pets",
	}

	for _, test := range palindromeTests {
		fmt.Printf("Is %q a palindrome? %v\n", test, isPalindrome(test))
	}
}

