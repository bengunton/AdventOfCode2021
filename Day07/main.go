package main

import (
	"fmt"
	"sort"

	"github.com/bengunton/AdventOfCode2021/util"
)

func main() {
	rows := util.ReadFileArg()
	positions := util.ParseInts(rows[0])

	findCheapestPointV2(positions)
}

func findCheapestPoint(positions []int) {
	midpoint := findMedian(positions)
	fmt.Println("Midpoint:", midpoint)

	fuelCost := calculateFuelCost(midpoint, positions)
	fmt.Println("Cost:", fuelCost)
}

func findCheapestPointV2(positions []int) {
	midpoint, otherMidpoint := findMeans(positions)
	fmt.Println("Midpoint:", midpoint)
	fmt.Println("Midpoint2:", otherMidpoint)

	fuelCost := calculateFuelCostV2(midpoint, positions)
	fuelCost2 := calculateFuelCostV2(otherMidpoint, positions)
	fmt.Println("Cost:", fuelCost)
	fmt.Println("Cost2:", fuelCost2)
}

func calculateFuelCost(midpoint int, positions []int) int {
	cost := 0

	for _, pos := range positions {
		cost += abs(pos - midpoint)
	}

	return cost
}

func calculateFuelCostV2(midpoint int, positions []int) int {
	cost := 0

	for _, pos := range positions {
		cost += costForDistanceV2(abs(pos - midpoint))
	}

	return cost
}

func findMedian(positions []int) int {
	sort.Ints(positions)

	midpoint := positions[len(positions)/2]
	return midpoint
}

func findMeans(positions []int) (int, int) {
	sum := 0
	for _, pos := range positions {
		sum += pos
	}

	mid := sum / len(positions)
	return mid, mid + 1
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func costForDistanceV2(num int) int {
	return (num * (num + 1)) / 2
}
