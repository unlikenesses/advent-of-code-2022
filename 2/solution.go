package main

import (
	"fmt"

	"github.com/unlikenesses/utils"
)

var values = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func main() {
	lines := utils.ReadInput()
	r1 := partOne(lines)
	fmt.Println(r1)
	r2 := partTwo(lines)
	fmt.Println(r2)
}

func partOne(lines []string) int {
	total := 0
	for _, line := range lines {
		opponent := string(line[0])
		me := string(line[2])
		score := values[me] + outcome1(opponent, me)
		total += score
	}

	return total
}

func partTwo(lines []string) int {
	var outcomeValues = map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	total := 0
	for _, line := range lines {
		opponent := string(line[0])
		outcome := string(line[2])
		score := outcomeValues[outcome] + values[outcome2(opponent, outcome)]
		total += score
	}
	return total
}

func outcome1(x string, y string) int {
	switch x {
	case "A":
		if y == "X" {
			return 3
		}
		if y == "Y" {
			return 6
		}
		if y == "Z" {
			return 0
		}
	case "B":
		if y == "X" {
			return 0
		}
		if y == "Y" {
			return 3
		}
		if y == "Z" {
			return 6
		}
	case "C":
		if y == "X" {
			return 6
		}
		if y == "Y" {
			return 0
		}
		if y == "Z" {
			return 3
		}
	}
	return 0
}

func outcome2(x string, y string) string {
	switch y {
	case "X":
		if x == "A" {
			return "Z"
		}
		if x == "B" {
			return "X"
		}
		if x == "C" {
			return "Y"
		}
	case "Y":
		if x == "A" {
			return "X"
		}
		if x == "B" {
			return "Y"
		}
		if x == "C" {
			return "Z"
		}
	case "Z":
		if x == "A" {
			return "Y"
		}
		if x == "B" {
			return "Z"
		}
		if x == "C" {
			return "X"
		}
	}
	return ""
}
