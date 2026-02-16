package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"ru.prostoyartemka.mystm32/compiler"
)

func main() {
	var output []byte

	args := os.Args

	if len(args) == 1 {

		fmt.Println("Usage:")
		fmt.Println("\tcompiler <path_to_input> <path_to_output>")

		return
	}

	if len(args) != 3 {
		return
	}

	inputPath := args[1]
	outputPath := args[2]

	file, err := os.Open(inputPath)

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

	file, err = os.Create(outputPath)

	if err != nil {
		log.Fatal(err)
	}

	file.Write(output)

	file.Close()

	fmt.Println("Compiled succsefuly to", outputPath)
}
