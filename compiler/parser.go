package compiler

import (
	"bytes"
	"encoding/binary"
	"log"
	"strconv"
	"strings"
)

var LABELS map[string]int = map[string]int{}
var GLOBAL_POINTER int

func intToBytes(i int32) []byte {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, i)

	if err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}

func GetInstructionsWithSuffix(token Token) ([]Instruction, string) {

	for name, instruction := range INSTRUCTIONS {

		if strings.HasPrefix(token.Value, name) {
			withoutPrefix, _ := strings.CutPrefix(token.Value, name)

			if strings.TrimSpace(withoutPrefix) == "" {
				return instruction, "AL"
			}

			for _, suffix := range SUFFXIES {
				if withoutPrefix == suffix {
					return instruction, suffix
				}
			}
		}

	}

	return nil, "AL"

}

func GetInstructionByLen(instructions []Instruction, length int) Instruction {

	for _, i := range instructions {
		if len(i.Args) == length {
			return i
		}
	}

	return Instruction{}
}

func ParseArguments(instruction Instruction, tokens []Token) []byte {

	var bytes []byte

	for index, arg := range tokens {

		argExpect := instruction.Args[index]
		token := tokens[index]
		num, _ := strconv.Atoi(token.Value)

		if arg.Type == TOKEN_REGISTER {

			bytes = append(bytes, TYPE_REGISTER)
			bytes = append(bytes, byte(num))

			if argExpect == TOKEN_OPERAND {
				bytes = append(bytes, 0, 0, 0) // Empty space for packing equality
			}
		}

		if arg.Type == TOKEN_CONSTANT {

			bytes = append(bytes, TYPE_CONSTANT)
			bytes = append(bytes, intToBytes(int32(num))...)

		}

		if arg.Type == TOKEN_ARGUMENT_LABEL {

			bytes = append(bytes, TYPE_LABEL)
			bytes = append(bytes, intToBytes(int32(LABELS[arg.Value]))...)

		}

		// WIP: Other types of args

	}

	return bytes
}

func GetArgumentSize(token Token, index int, instruction Instruction) int {

	argExpect := instruction.Args[index]
	result := 0

	if token.Type == TOKEN_REGISTER {
		result = 2

		if argExpect == TOKEN_OPERAND {
			return 5
		}

		return result
	}

	if token.Type == TOKEN_CONSTANT {

		return 5

	}

	return 0
}

func GetInstructionSize(tokens []Token, index int) int {
	var token = tokens[index]

	var args []Token

	for {
		index++

		if index >= len(tokens) {
			break
		}

		if tokens[index].Type != TOKEN_INSTRUCTION {
			args = append(args, tokens[index])

			continue
		}

		break
	}

	instructions, suffix := GetInstructionsWithSuffix(token)

	instruction := GetInstructionByLen(instructions, len(args))

	instruction.Suffix = suffix

	var argsSize = 0

	for i, arg := range args {
		argsSize += GetArgumentSize(arg, i, instruction)
	}

	return 2 + len(instruction.Prefix) + argsSize
}

func ParseLabels(tokens []Token) {

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if token.Type == TOKEN_INSTRUCTION {

			GLOBAL_POINTER += GetInstructionSize(tokens, i)

			continue
		}

		if token.Type == TOKEN_LABEL {
			LABELS[token.Value] = GLOBAL_POINTER
		}
	}
}

func Parse(tokens []Token) []byte {

	var bytes []byte

	for i := 0; i < len(tokens); i++ {
		var token = tokens[i]

		if token.Type == TOKEN_LABEL {
			continue
		}

		var instructionToken = token
		var args []Token

		for {
			i++

			if i >= len(tokens) {
				break
			}

			if tokens[i].Type != TOKEN_INSTRUCTION {
				args = append(args, tokens[i])

				continue
			}

			i--

			break
		}

		instructions, suffix := GetInstructionsWithSuffix(instructionToken)

		instruction := GetInstructionByLen(instructions, len(args))
		instruction.Suffix = suffix

		bytes = append(bytes, instruction.Byte)
		bytes = append(bytes, SUFFIX_BYTES[suffix])
		bytes = append(bytes, instruction.Prefix...)
		bytes = append(bytes, ParseArguments(instruction, args)...)

	}

	return bytes
}
