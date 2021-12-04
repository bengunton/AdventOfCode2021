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

type Submarine struct {
	loc Point
	aim int
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

	start := Submarine{Point{0, 0}, 0}

	instructions := ReadInstructions(file)
	end := FollowAimedCourse(start, instructions)

	fmt.Printf("Product is %d", end.loc.depth*end.loc.distance)
}

func FollowSimpleCourse(start Point, instructions []string) Point {
	end := start
	for _, row := range instructions {
		instruction := parseInstruction(row)
		switch instruction.command {
		case "up":
			end.depth -= instruction.amount
		case "down":
			end.depth += instruction.amount
		case "forward":
			end.distance += instruction.amount
		}
	}

	return end
}

func FollowAimedCourse(sub Submarine, instructions []string) Submarine {
	for _, row := range instructions {
		instruction := parseInstruction(row)
		switch instruction.command {
		case "up":
			sub.aim -= instruction.amount
		case "down":
			sub.aim += instruction.amount
		case "forward":
			sub.loc.distance += instruction.amount
			sub.loc.depth += sub.aim * instruction.amount
		}
	}

	return sub
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
