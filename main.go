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

	lines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")

	for lineIndex, line := range lines {
		words := strings.Fields(line)

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

		for i := 0; i+1 < len(words); i++ {
			if strings.EqualFold(words[i], "a") {
				nextWord := strings.TrimLeft(words[i+1], ".,!?;:")

				if len(nextWord) > 0 {
					firstLetter := strings.ToLower(nextWord[:1])

					if strings.Contains("aeiouh", firstLetter) {
						if words[i] == "A" {
							words[i] = "An"
						} else {
							words[i] = "an"
						}
					}
				}
			}
		}

		lineText := fixPunctuation(strings.Join(words, " "))
		lines[lineIndex] = fixApostrophes(lineText)
	}

	result := strings.Join(lines, "\n")

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

func fixApostrophes(text string) string {
	chars := []rune(text)

	for i := 0; i < len(chars); i++ {
		opening := chars[i]
		var closing rune

		switch opening {
		case '\'':
			closing = '\''
		case '‘':
			closing = '’'
		default:
			continue
		}

		for j := i + 1; j < len(chars); j++ {
			if chars[j] == closing || (opening == '‘' && chars[j] == '‘') {
				start := i + 1
				for start < j && unicode.IsSpace(chars[start]) {
					start++
				}

				end := j
				for end > start && unicode.IsSpace(chars[end-1]) {
					end--
				}

				fixed := make([]rune, 0, len(chars))
				fixed = append(fixed, chars[:i+1]...)
				fixed = append(fixed, chars[start:end]...)

				if opening == '‘' {
					fixed = append(fixed, '’')
				} else {
					fixed = append(fixed, closing)
				}

				fixed = append(fixed, chars[j+1:]...)
				chars = fixed
				i = i + 1 + (end - start)
				break
			}
		}
	}

	return string(chars)
}
