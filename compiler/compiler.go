package compiler

import "strings"

func PreCompile(line string) {

	if len(strings.TrimSpace(line)) == 0 {
		return
	}

	words := GetWords(line)
	tokens := GetTokens(words)

	ParseLabels(tokens)
}

func Compile(line string) []byte {

	if len(strings.TrimSpace(line)) == 0 {
		return []byte{}
	}

	words := GetWords(line)
	tokens := GetTokens(words)

	return Parse(tokens)
}
