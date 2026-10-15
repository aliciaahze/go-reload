package main

import (
	"fmt"
	"os"
	"strings"
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

	text := string(content)
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		if i == 0 {
			continue
		}

		switch words[i] {
		case "(up)":
			words[i-1] = strings.ToUpper(words[i-1])
		case "(low)":
			words[i-1] = strings.ToLower(words[i-1])
		case "(cap)":
			word := strings.ToLower(words[i-1])
			if len(word) > 0 {
				words[i-1] = strings.ToUpper(word[:1]) + word[1:]
			}
		default:
			continue
		}

		words = append(words[:i], words[i+1:]...)
		i--
	}

	result := strings.Join(words, " ")

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}

	fmt.Println("File transformed successfully!")
}
