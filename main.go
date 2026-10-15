package main

import (
	"fmt"
	"os"
	"strconv"
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

		if marker != "(up)" && marker != "(low)" && marker != "(cap)" {
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
			}
		}

		words = append(words[:i], words[i+removeCount:]...)
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
