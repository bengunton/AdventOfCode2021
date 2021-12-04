package main

import (
	"bufio"
	"fmt"
	"io"
)

type Point struct {
	distance int
	depth    int
}

func main() {
	fmt.Println("vim-go")
}

func ReadInstructions(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	var instructions []string

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	return instructions
}
