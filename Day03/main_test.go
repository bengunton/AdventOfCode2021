package main

import (
	"os"
	"testing"

	"github.com/bengunton/AdventOfCode2021/util"
)

func TestIsSetAtOffset(t *testing.T) {
	testCases := []struct {
		reading  Reading
		offset   int
		expected bool
	}{
		{1, 0, true},
		{1, 1, false},
		{1, 2, false},

		{2, 0, false},
		{2, 1, true},
		{2, 2, false},

		{3, 0, true},
		{3, 1, true},
		{3, 2, false},
	}

	for _, test := range testCases {
		actual := IsSetAtOffset(test.reading, test.offset)
		if actual != test.expected {
			t.Errorf("Expected IsSetAtOffset(%012b, %d) = %t", test.reading, test.offset,
				test.expected)
		}
	}
}

func TestCountOnesAtOffset(t *testing.T) {
	readings := []Reading{
		0b0000000000000010,
		0b0000000000000000,
		0b0000000000000010,
		0b0000000000000010,
	}

	actual := CountOnesAtOffset(readings, 1)
	if actual != 3 {
		t.Errorf("Expected 3 ones at offset 1 of readings, actual %d", actual)
	}
}

func TestReadGammaRate(t *testing.T) {
	input, err := os.Open("sample")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	inputs := util.ReadFile(input)
	bitLength := len(inputs[0])
	readings := ParseContents(inputs, bitLength)

	actual := ReadGammaRate(readings, bitLength)
	expected := Reading(22)

	if actual != expected {
		t.Errorf("Gamma rate expected %d, actual %d", expected, actual)
	}
}
