package main

import (
	"fmt"
	"strconv"

	"github.com/bengunton/AdventOfCode2021/util"
)

type Reading uint16

const ReadingBase = 2

func main() {
	file := util.OpenFileFromArgs()
	defer file.Close()

	rows := util.ReadFile(file)
	bitLength := len(rows[0])
	readings := ParseContents(rows, bitLength)

	gamma := ReadGammaRate(readings, bitLength)
	epsilon := CalculateEpsilon(gamma, bitLength)

	fmt.Printf("Gamma    %016b %d\n", gamma, gamma)
	fmt.Printf("Epsilion %016b %d\n", epsilon, epsilon)
	product := uint64(gamma) * uint64(epsilon)
	fmt.Printf("Product is %d\n", product)

	oxygen := CalculateOxygenRating(readings, bitLength)
	co2 := CalculateCO2Rating(readings, bitLength)

	fmt.Printf("Oxygen %016b %d\n", oxygen, oxygen)
	fmt.Printf("CO2    %016b %d\n", co2, co2)
	product = uint64(oxygen) * uint64(co2)
	fmt.Printf("Product is %d\n", product)
}

func ParseContents(rows []string, readingBits int) []Reading {
	var readings []Reading
	for _, row := range rows {
		value, err := strconv.ParseUint(row, ReadingBase, readingBits)
		if err != nil {
			panic("Unable to parse value")
		}
		readings = append(readings, Reading(value))
	}

	return readings
}

func ReadGammaRate(readings []Reading, readingBits int) Reading {
	gamma := Reading(0)
	for offset := 0; offset < readingBits; offset++ {
		ones := CountOnesAtOffset(readings, offset)
		if ones > len(readings)/2 {
			gamma += 1 << offset
		}
	}

	return gamma
}

func CalculateEpsilon(gamma Reading, readingBits int) Reading {
	mask := (Reading(1) << readingBits) - 1
	return gamma ^ mask
}

func CalculateOxygenRating(readings []Reading, readingBits int) Reading {
	for offset := readingBits - 1; offset >= 0; offset-- {
		ones := CountOnesAtOffset(readings, offset)
		if float32(ones) >= float32(len(readings))/2 {
			readings, _ = splitByOffset(readings, offset)
		} else {
			_, readings = splitByOffset(readings, offset)
		}
		if len(readings) == 1 {
			return readings[0]
		}
	}

	return 0
}

func CalculateCO2Rating(readings []Reading, readingBits int) Reading {
	for offset := readingBits - 1; offset >= 0; offset-- {
		ones := CountOnesAtOffset(readings, offset)
		if float32(ones) < float32(len(readings))/2 {
			readings, _ = splitByOffset(readings, offset)
		} else {
			_, readings = splitByOffset(readings, offset)
		}
		if len(readings) == 1 {
			return readings[0]
		}
	}

	return 0
}

func splitByOffset(readings []Reading, offset int) ([]Reading, []Reading) {
	mask := Reading(1) << offset
	var ones []Reading
	var zeros []Reading

	for _, reading := range readings {
		if mask&reading > 0 {
			ones = append(ones, reading)
		} else {
			zeros = append(zeros, reading)
		}
	}
	return ones, zeros
}

func CountOnesAtOffset(readings []Reading, offset int) int {
	ones := 0

	for _, reading := range readings {
		if IsSetAtOffset(reading, offset) {
			ones++
		}
	}
	return ones
}

func IsSetAtOffset(reading Reading, offset int) bool {
	column := uint16(uint16(1) << offset)
	return (uint16(reading) & column) > 0
}
