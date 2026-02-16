package compiler

import (
	"maps"
	"slices"
)

const (
	TYPE_REGISTER = iota
	TYPE_CONSTANT
	TYPE_REGISTER_SHIFT
	TYPE_MEM_LOAD
	TYPE_LABEL
)

var SUFFIX_BYTES map[string]byte = map[string]byte{

	"EQ": 0, "NE": 1, "CS": 2, "HS": 3, "CC": 4, "LO": 5, "MI": 6, "PL": 7, "VS": 8, "VC": 9, "HI": 10, "LS": 11, "GE": 12, "LT": 13, "GT": 14, "LE": 15, "AL": 16, "S": 17,
}

var SUFFXIES = slices.Collect(maps.Keys(SUFFIX_BYTES))

var INSTRUCTIONS map[string][]Instruction = map[string][]Instruction{
	"ADD": {
		{Byte: 1, Args: []int{TOKEN_REGISTER, TOKEN_OPERAND}, Prefix: []byte{0, 0, 0}},
		{Byte: 1, Args: []int{TOKEN_REGISTER, TOKEN_REGISTER, TOKEN_OPERAND}, Prefix: []byte{1}},
	},

	"SUB": {
		{Byte: 2, Args: []int{TOKEN_REGISTER, TOKEN_OPERAND}, Prefix: []byte{0, 0, 0}},
		{Byte: 2, Args: []int{TOKEN_REGISTER, TOKEN_REGISTER, TOKEN_OPERAND}, Prefix: []byte{1}},
	},

	"MUL": {
		{Byte: 3, Args: []int{TOKEN_REGISTER, TOKEN_REGISTER}, Prefix: []byte{0, 0, 0}},
		{Byte: 3, Args: []int{TOKEN_REGISTER, TOKEN_REGISTER, TOKEN_REGISTER}, Prefix: []byte{1}},
	},

	"B": {
		{Byte: 4, Args: []int{TOKEN_LABEL}},
	},

	"MOV": {
		{Byte: 5, Args: []int{TOKEN_REGISTER, TOKEN_OPERAND}},
	},

	"SDIV": {
		{Byte: 6, Args: []int{TOKEN_REGISTER, TOKEN_REGISTER}, Prefix: []byte{0}},
		{Byte: 6, Args: []int{TOKEN_REGISTER, TOKEN_REGISTER, TOKEN_REGISTER}, Prefix: []byte{1}},
	},
}

type Instruction struct {
	Byte byte

	Args   []int
	Prefix []byte

	Suffix string
}
