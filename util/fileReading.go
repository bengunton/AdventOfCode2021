package util

import (
	"bufio"
	"io"
	"os"
)

func ReadFile(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func OpenFileFromArgs() *os.File {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	return file
}

func ReadFileArg() []string {
	file := OpenFileFromArgs()
	defer file.Close()

	return ReadFile(file)
}
