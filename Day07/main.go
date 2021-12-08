package main

import (
	"fmt"
	"sort"

	"github.com/bengunton/AdventOfCode2021/util"
)

func main() {
	rows := util.ReadFileArg()
	positions := util.ParseInts(rows[0])

	findCheapestPoint(positions)
}

func findCheapestPoint(positions []int) {
	midpoint := findMedian(positions)
	fmt.Println("Midpoint:", midpoint)

	fuelCost := calculateFuelCost(midpoint, positions)
	fmt.Println("Cost:", fuelCost)
}

func calculateFuelCost(midpoint int, positions []int) int {
	cost := 0

	for _, pos := range positions {
		cost += abs(pos - midpoint)
	}

	return cost
}

func findMedian(positions []int) int {
	sort.Ints(positions)

	midpoint := positions[len(positions)/2]
	return midpoint
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
