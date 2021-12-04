package main

import (
	"os"
	"testing"

	"github.com/bengunton/AdventOfCode2021/util"
)

func TestFollowSimpleCourse(t *testing.T) {
	start := Point{distance: 0, depth: 0}

	input, err := os.Open("sample")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	instructions := util.ReadFile(input)

	end := FollowSimpleCourse(start, instructions)
	product := end.distance * end.depth
	if product != 150 {
		t.Errorf("Expected product of 150 , got %d", product)
	}
}

func TestFollowAimedCourse(t *testing.T) {
	start := Submarine{loc: Point{distance: 0, depth: 0}, aim: 0}

	input, err := os.Open("sample")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	instructions := util.ReadFile(input)

	end := FollowAimedCourse(start, instructions)
	product := end.loc.distance * end.loc.depth
	if product != 900 {
		t.Errorf("Expected product of 900, got %d", product)
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
