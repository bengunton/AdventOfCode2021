package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/bengunton/AdventOfCode2021/util"
)

type Line struct {
	start Point
	end   Point
}

type Point struct {
	x int
	y int
}

type Grid [1000][1000]int

func main() {
	rows := util.ReadFileArg()
	lines := parseLines(rows)

	grid := new(Grid)

	DrawLinesComplex(lines, grid)
	count := grid.CountAtLeastTwos()
	fmt.Printf("Squares with at least 2: %d", count)
}

func (g *Grid) CountAtLeastTwos() int {
	count := 0
	for _, row := range g {
		for _, square := range row {
			if square >= 2 {
				count++
			}
		}
	}
	return count
}

func DrawLinesSimple(lines []Line, g *Grid) {
	for _, line := range lines {
		if line.start.x != line.end.x && line.start.y != line.end.y {
			continue
		}

		if line.start.x == line.end.x {
			if line.start.y < line.end.y {
				for i := line.start.y; i <= line.end.y; i++ {
					g[i][line.start.x]++
				}
			} else {
				for i := line.start.y; i >= line.end.y; i-- {
					g[i][line.start.x]++
				}
			}

		}

		if line.start.y == line.end.y {
			if line.start.x < line.end.x {
				for i := line.start.x; i <= line.end.x; i++ {
					g[line.start.y][i]++
				}
			} else {
				for i := line.start.x; i >= line.end.x; i-- {
					g[line.start.y][i]++
				}
			}
		}
	}
}

func DrawLinesComplex(lines []Line, g *Grid) {
	DrawLinesSimple(lines, g)

	for _, line := range lines {
		if line.start.x == line.end.x || line.start.y == line.end.y {
			continue
		}

		if line.start.x < line.end.x {
			if line.start.y < line.end.y {
				for i := 0; i <= line.end.x-line.start.x; i++ {
					g[line.start.y+i][line.start.x+i]++
				}
			} else {
				for i := 0; i <= line.end.x-line.start.x; i++ {
					g[line.start.y-i][line.start.x+i]++
				}
			}
		} else {
			if line.start.y < line.end.y {
				for i := 0; i <= line.start.x-line.end.x; i++ {
					g[line.start.y+i][line.start.x-i]++
				}
			} else {
				for i := 0; i <= line.start.x-line.end.x; i++ {
					g[line.start.y-i][line.start.x-i]++
				}
			}
		}
	}
}

func parseLines(rows []string) []Line {
	var lines []Line
	for _, row := range rows {
		line, err := parseLine(row)
		if err != nil {
			panic(err)
		}
		lines = append(lines, line)
	}
	return lines
}

const lineRegex = `^(\d*,\d*) -> (\d*,\d*)$`

func parseLine(row string) (Line, error) {
	regex := regexp.MustCompile(lineRegex)
	matches := regex.FindStringSubmatch(row)

	if len(matches) != 3 {
		return Line{}, errors.New(fmt.Sprintf("Not enough matches in %s, got %v", row, matches))
	}
	start, err := parsePoint(matches[1])
	if err != nil {
		return Line{}, err
	}
	end, err := parsePoint(matches[2])
	if err != nil {
		return Line{}, err
	}
	return Line{start, end}, nil
}

func parsePoint(input string) (Point, error) {
	coords := strings.Split(input, ",")
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		return Point{}, err
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		return Point{}, err
	}
	return Point{x, y}, nil
}
