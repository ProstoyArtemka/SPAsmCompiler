package compiler

import (
	"slices"
	"strings"
)

var BLANK_SPACES = []rune{' ', '\t', '\n', '\r'}

func GetWords(line string) []string {
	var words []string
	var word string = ""

	for _, char := range line {

		if char == ';' {
			break
		}

		if slices.Contains(BLANK_SPACES, char) {
			if word != "" {
				words = append(words, word)

				word = ""
			}

			continue
		}

		if char == ',' {
			words = append(words, word)
			word = ""

			words = append(words, ",")

			continue
		}

		word += string(char)
	}

	if strings.TrimSpace(word) != "" {
		words = append(words, word)
	}

	return words
}
