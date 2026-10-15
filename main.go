package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	words := strings.Fields(string(content))

	for i := 0; i < len(words); i++ {
		marker := words[i]
		count := 1
		removeCount := 1

		if (marker == "(up," || marker == "(low," || marker == "(cap,") &&
			i+1 < len(words) {
			n, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
			if err == nil && n > 0 {
				marker = strings.TrimSuffix(marker, ",") + ")"
				count = n
				removeCount = 2
			}
		}

		if marker != "(up)" && marker != "(low)" &&
			marker != "(cap)" && marker != "(hex)" &&
			marker != "(bin)" {
			continue
		}

		start := i - count
		if start < 0 {
			start = 0
		}

		for j := start; j < i; j++ {
			switch marker {
			case "(up)":
				words[j] = strings.ToUpper(words[j])
			case "(low)":
				words[j] = strings.ToLower(words[j])
			case "(cap)":
				word := strings.ToLower(words[j])
				if len(word) > 0 {
					words[j] = strings.ToUpper(word[:1]) + word[1:]
				}
			case "(hex)":
				value, err := strconv.ParseInt(words[j], 16, 64)
				if err == nil {
					words[j] = strconv.FormatInt(value, 10)
				}

			case "(bin)":
				value, err := strconv.ParseInt(words[j], 2, 64)
				if err == nil {
					words[j] = strconv.FormatInt(value, 10)
				}
			}
		}

		words = append(words[:i], words[i+removeCount:]...)
		i--
	}

	result := fixPunctuation(strings.Join(words, " "))

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}

	fmt.Println("File transformed successfully!")
}

func fixPunctuation(text string) string {
	punctuation := ".,!?;:"
	input := []rune(strings.TrimSpace(text))
	output := []rune{}

	for i, char := range input {
		if unicode.IsSpace(char) {
			if len(output) > 0 &&
				!unicode.IsSpace(output[len(output)-1]) {
				output = append(output, ' ')
			}
			continue
		}

		if strings.ContainsRune(punctuation, char) {
			for len(output) > 0 &&
				unicode.IsSpace(output[len(output)-1]) {
				output = output[:len(output)-1]
			}

			output = append(output, char)

			if i+1 < len(input) &&
				!unicode.IsSpace(input[i+1]) &&
				!strings.ContainsRune(punctuation, input[i+1]) {
				output = append(output, ' ')
			}
			continue
		}

		output = append(output, char)
	}

	return strings.TrimSpace(string(output))
}
