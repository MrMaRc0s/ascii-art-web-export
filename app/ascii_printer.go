package main

import (
	"strings"
)

// generateAsciiArt generates the ASCII art for the input text using the chosen banner
func GenerateAsciiArt(inputText, banner string) (string, error) {
	lines, err := loadBanner(banner)
	if err != nil {
		return "", err
	}

	// Split input text by literal "\n" for handling line breaks
	sepText := strings.Split(inputText, "\\n")
	var result strings.Builder
	printAsciiArtRecursive(sepText, lines, &result) // Recursively build the ASCII art
	return result.String(), nil
}

// PrintAsciiArtRecursive generates ASCII art for each sentence recursively
func printAsciiArtRecursive(sentences []string, textFile []string, result *strings.Builder) {
	// Base case: If no sentences left, return
	if len(sentences) == 0 {
		return
	}

	// Process the first sentence
	if sentences[0] != "" {
		printSentenceAsciiRecursive(sentences[0], textFile, 1, result)
		result.WriteString("\n") // Add a newline after each sentence
	} else {
		result.WriteString("\n") // Empty line for "\n" in input
	}

	// Recursive call to process remaining sentences
	printAsciiArtRecursive(sentences[1:], textFile, result)
}

// printSentenceAsciiRecursive prints each line of ASCII art for a word recursively
func printSentenceAsciiRecursive(word string, textFile []string, h int, result *strings.Builder) {
	// Base case: If line height exceeds the banner height (8), return
	if h > 8 {
		return
	}

	// Print the ASCII art line by line for each character
	for i := 0; i < len(word); i++ {
		for lineIndex, line := range textFile {
			if lineIndex == (int(word[i])-32)*9+h {
				result.WriteString(line)
			}
		}
	}
	result.WriteString("\n") // Newline after each row of ASCII art for a word

	// Recursive call to print the next line height
	printSentenceAsciiRecursive(word, textFile, h+1, result)
}
