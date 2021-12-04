package main

import (
	"os"
	"testing"
)

/*
func TestFollowCourse(t *testing.T) {
	start := Point{distance: 0, depth: 0}

	input, err := os.Open("sample")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	instructions := ReadInstructions(input)

	end := FollowCourse(start, instructions)
	product := end.Width * end.Height
	if product != 150 {
		t.Errorf("Expected product of 150 , got %d", product)
	}
}
*/

func TestReadInstructions(t *testing.T) {
	sample, err := os.Open("sample")
	if err != nil {
		panic(err)
	}
	defer sample.Close()

	expected := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	actual := ReadInstructions(sample)

	if len(actual) != len(expected) {
		t.Fatalf("Array length differed, expected %d, actual %d)", len(expected), len(actual))
	}
	for i, entry := range expected {
		if actual[i] != entry {
			t.Errorf("Expected entry %s, got %s", expected, actual[i])
		}
	}
}

func TestParseInstruction(t *testing.T) {
	actual := parseInstruction("down 5")
	if actual.command != "down" {
		t.Errorf("Failed to parse command")
	}
	if actual.amount != 5 {
		t.Errorf("Failed to parse amount")
	}
}
