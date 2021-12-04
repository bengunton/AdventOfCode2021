package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bengunton/AdventOfCode2021/util"
)

type BingoCard struct {
	card     [5][5]int
	markings [5][5]bool
}

type BingoDraws []int

func main() {
	file := util.OpenFileFromArgs()
	defer file.Close()

	rows := util.ReadFile(file)

	draws := ParseBingoDraws(rows[0])
	cards := ParseBingoCards(rows[1:])

	card, i, last := Play(draws, cards)

	fmt.Printf("Winner is %d\n", i+1)
	sum := card.sumUnmarked()
	fmt.Printf("Sum is %d\n", sum)
	fmt.Printf("Last called * Sum = %d\n", last*sum)
}

func Play(draws BingoDraws, cards []BingoCard) (BingoCard, int, int) {
	for _, draw := range draws {
		for i, _ := range cards {
			cards[i].mark(draw)
			if cards[i].checkWin() {
				return cards[i], i, draw
			}
		}
	}

	return BingoCard{}, -1, -1
}

func (c *BingoCard) mark(draw int) {
	for i, row := range c.card {
		for j, num := range row {
			if num == draw {
				c.markings[i][j] = true
			}
		}
	}
}

func (c *BingoCard) checkWin() bool {
	for i, row := range c.markings {
		// rows
		win := true
		for _, marked := range row {
			if !marked {
				win = false
			}
		}

		if win {
			return true
		}

		// columns
		win = true
		for _, row2 := range c.markings {
			if !row2[i] {
				win = false
			}
		}

		if win {
			return true
		}
	}
	return false
}

func (c *BingoCard) sumUnmarked() int {
	sum := 0

	for i, row := range c.markings {
		for j, marking := range row {
			if !marking {
				sum += c.card[i][j]
			}
		}
	}

	return sum
}

func ParseBingoDraws(draws string) BingoDraws {
	drawsArr := strings.Split(draws, ",")

	return parseNumsFromStrings(drawsArr)
}

func ParseBingoCards(rows []string) []BingoCard {
	bingoRow := 0
	var bingoCard BingoCard
	var bingoCards []BingoCard

	for _, row := range rows[1:] {
		if row == "" {
			bingoCards = append(bingoCards, bingoCard)
			bingoCard = BingoCard{}
			bingoRow = 0
			continue
		}

		numRow := strings.Fields(row)
		rowSlice := parseNumsFromStrings(numRow)
		copy(bingoCard.card[bingoRow][0:5], rowSlice[:])
		bingoRow++
	}

	bingoCards = append(bingoCards, bingoCard)

	return bingoCards
}

func parseNumsFromStrings(strArr []string) []int {
	var numbers []int

	for _, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, num)
	}

	return numbers
}
