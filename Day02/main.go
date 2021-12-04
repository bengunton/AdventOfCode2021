package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	distance int
	depth    int
}

type Instruction struct {
	command string
	amount  int
}

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	start := Point{0, 0}

	instructions := ReadInstructions(file)
	FollowCourse(start, instructions)
}

func FollowCourse(start Point, instructions []string) Point {
	for _, instruction := range instructions {
		fmt.Println(instruction)
	}

	return start
}

func parseInstruction(instruction string) Instruction {
	parts := strings.Split(instruction, " ")
	if len(parts) != 2 {
		panic("Bad instruction")
	}
	amount, _ := strconv.Atoi(parts[1])
	return Instruction{parts[0], amount}
}

func ReadInstructions(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	var instructions []string

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	return instructions
}
