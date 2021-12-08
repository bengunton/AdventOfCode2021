package util

import (
	"strconv"
	"strings"
)

func ParseInts(input string) []int {
	data := strings.Split(input, ",")
	var result []int

	for _, datum := range data {
		num, _ := strconv.Atoi(datum)
		result = append(result, num)
	}

	return result
}
