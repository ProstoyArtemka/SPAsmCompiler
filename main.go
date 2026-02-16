package main

import (
	"bufio"
	"log"
	"os"

	"ru.prostoyartemka.mystm32/compiler"
)

const INPUT_FILE = "code/input.s"
const OUTPUT_FILE = "code/out.bin"

func main() {
	var output []byte

	file, err := os.Open(INPUT_FILE)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Compiling

	for _, line := range lines {
		compiler.PreCompile(line)
	}

	for _, line := range lines {
		res := compiler.Compile(line)

		output = append(output, res...)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Write output

	file, err = os.Create(OUTPUT_FILE)

	if err != nil {
		log.Fatal(err)
	}

	file.Write(output)

	file.Close()
}
