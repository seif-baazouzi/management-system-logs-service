package utils

import (
	"strconv"
	"strings"
)

func GetMonthDays(m string) uint8 {
	splittedList := strings.Split(m, "-")
	month, _ := strconv.Atoi(splittedList[0])
	year, _ := strconv.Atoi(splittedList[1])

	switch month {
	case 2:
		if year%4 == 0 {
			return 29
		}
		return 28

	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	default:
		return 30
	}
}
