package compiler

import (
	"maps"
	"regexp"
	"slices"
	"strings"
)

const (
	TOKEN_INSTRUCTION = iota

	TOKEN_REGISTER
	TOKEN_CONSTANT
	TOKEN_OPERAND
	TOKEN_LABEL
	TOKEN_ARGUMENT_LABEL

	NEWLINE
)

var INSTRUCTION_NAMES = []string{"ADD"}

var CONSTANT_REGEX = regexp.MustCompile(`^#-?\d+`)
var REGISTER_REGEX = regexp.MustCompile(`^r[0-9]{1,2}`)
var LABEL_REGEX = regexp.MustCompile(`[a-zA-Z_-]+[0-9a-zA-Z_-]*:$`)

type Token struct {
	Type int

	Value string
}

func GetTokens(words []string) []Token {

	var tokens []Token
	var labels = slices.Collect(maps.Keys(LABELS))

	for i := 0; i < len(words); i++ {
		word := words[i]

		nextIsInstruction := true
		if len(tokens) > 0 {
			if tokens[len(tokens)-1].Type == TOKEN_INSTRUCTION {
				nextIsInstruction = false
			}
		}

		if word == "," {
			nextIsInstruction = false

			if i+1 >= len(words) {
				break
			}

			i++
			word = words[i]
		}

		if LABEL_REGEX.MatchString(word) {
			label_name, _ := strings.CutSuffix(word, ":")

			tokens = append(tokens, Token{
				Type: TOKEN_LABEL, Value: label_name,
			})

			nextIsInstruction = true

			continue
		}

		if slices.Contains(labels, word) {

			tokens = append(tokens, Token{
				Type: TOKEN_ARGUMENT_LABEL, Value: word,
			})

			continue
		}

		if nextIsInstruction {
			tokens = append(tokens, Token{
				Type: TOKEN_INSTRUCTION, Value: word,
			})

			continue
		}

		if CONSTANT_REGEX.MatchString(word) {
			num, _ := strings.CutPrefix(word, "#")

			tokens = append(tokens, Token{
				Type: TOKEN_CONSTANT, Value: num,
			})

			continue
		}

		if REGISTER_REGEX.MatchString(word) {
			index, _ := strings.CutPrefix(word, "r")

			tokens = append(tokens, Token{
				Type: TOKEN_REGISTER, Value: index,
			})

			continue
		}
	}

	return tokens
}
