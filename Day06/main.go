package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bengunton/AdventOfCode2021/util"
)

const DAYS = 256

func main() {
	rows := util.ReadFileArg()
	fish := ParseFish(rows[0])
	fmt.Printf("%v\n", fish)

	// fish = Simulate(fish, DAYS)
	sum := SimulateEfficient(fish, DAYS)
	fmt.Println("Number of fish is", sum)
}

func SimulateEfficient(fishies []LanternFish, days int) int {
	tank := FillFishTank(fishies)

	for i := 0; i < days; i++ {
		tank.Age()
	}

	return tank.Count()
}

func Simulate(fishies []LanternFish, days int) []LanternFish {

	for i := 0; i < days; i++ {
		var newFishies []LanternFish

		for i, _ := range fishies {
			newFish := fishies[i].Age()
			newFishies = append(newFishies, newFish...)
		}

		fishies = append(fishies, newFishies...)
	}

	return fishies
}

func ParseFish(input string) []LanternFish {
	var fishes []LanternFish

	splitInput := strings.Split(input, ",")
	for _, part := range splitInput {
		fish, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}

		fishes = append(fishes, LanternFish(fish))
	}
	return fishes
}

func FillFishTank(fishies []LanternFish) FishTank {
	var tank FishTank

	for _, fish := range fishies {
		tank[fish]++
	}

	return tank
}
