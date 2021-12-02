package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Printf("Counted %d increases in %s", CountThreeWindowDepthIncreases(file), fileName)
}

func CountDepthIncreases(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)

	numIncreases := 0
	var prevVal int

	// first line
	scanner.Scan()
	currentVal := readLine(scanner)

	// remaining lines
	for scanner.Scan() {
		prevVal = currentVal
		currentVal = readLine(scanner)

		if currentVal > prevVal {
			numIncreases++
		}
	}

	return numIncreases
}

func CountThreeWindowDepthIncreases(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)

	numIncreases := 0
	var prevSum int
	var currSum int

	// first line
	scanner.Scan()
	reading1 := readLine(scanner)

	// second line
	scanner.Scan()
	reading2 := readLine(scanner)

	// third line
	scanner.Scan()
	reading3 := readLine(scanner)

	currSum = reading1 + reading2 + reading3
	// remaining lines
	for scanner.Scan() {
		prevSum = currSum

		reading1, reading2, reading3 = reading2, reading3, readLine(scanner)
		currSum = reading1 + reading2 + reading3

		if currSum > prevSum {
			numIncreases++
		}

	}

	return numIncreases
}

func readLine(scanner *bufio.Scanner) int {
	val, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	return val
}
