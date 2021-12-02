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

	fmt.Printf("Counted %d increases in %s", CountDepthIncreases(file), fileName)
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

func readLine(scanner *bufio.Scanner) int {
	val, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	return val
}
