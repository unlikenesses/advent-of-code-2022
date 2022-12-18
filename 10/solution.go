package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/unlikenesses/utils"
)

type AddCommand struct {
	toAdd      int
	cycleToAdd int
}

var lines []string

func main() {
	lines = utils.ReadInput()
	r1 := partOne()
	fmt.Println(r1)
	partTwo()
}

func partOne() int {
	x := 1
	numCycles := 220
	linePos := 0
	toAdd := 0
	signalStrengthSum := 0
	justAdded := false
	re := regexp.MustCompile(`addx (-?\d+)`)
	for c := 1; c <= numCycles; c++ {
		if c == 20 || c == 60 || c == 100 || c == 140 || c == 180 || c == 220 {
			signalStrengthSum += c * x
		}
		if justAdded {
			x += toAdd
			justAdded = false
			continue
		}
		command := lines[linePos]
		linePos++
		if linePos == len(lines) {
			linePos = 0
		}
		if command == "noop" {
			continue
		}
		matches := re.FindStringSubmatch(command)
		if len(matches) > 0 {
			toAdd, _ = strconv.Atoi(matches[1])
			justAdded = true
		} else {
			fmt.Println("No matches for line ", command)
		}
	}
	return signalStrengthSum
}

func partTwo() {
	var pixel string
	x := 1
	screenWidth := 40
	screenHeight := 6
	numCycles := screenWidth * screenHeight
	linePos := 0
	toAdd := 0
	justAdded := false
	cursor := 0
	row := 0
	re := regexp.MustCompile(`addx (-?\d+)`)
	for c := 0; c < numCycles; c++ {
		if c > 0 && c%screenWidth == 0 {
			fmt.Println()
			row = c / screenWidth
		}
		xPosInRow := x
		cursor = c - (row * screenWidth)
		if cursor >= xPosInRow-1 && cursor <= xPosInRow+1 {
			pixel = "#"
		} else {
			pixel = "."
		}
		fmt.Print(pixel)
		if justAdded {
			x += toAdd
			justAdded = false
			continue
		}
		command := lines[linePos]
		linePos++
		if linePos == len(lines) {
			linePos = 0
		}
		if command == "noop" {
			continue
		}
		matches := re.FindStringSubmatch(command)
		if len(matches) > 0 {
			toAdd, _ = strconv.Atoi(matches[1])
			justAdded = true
		} else {
			fmt.Println("No matches for line ", command)
		}
		cursor++
	}
}
