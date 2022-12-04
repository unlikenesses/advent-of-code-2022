package main

import (
	"fmt"
	"strings"

	"github.com/unlikenesses/utils"
)

type Rucksack struct {
	left  string
	right string
}

func main() {
	lines := utils.ReadInput()
	rucksacks := fillRucksacks(lines)
	r1 := partOne(rucksacks)
	fmt.Println(r1)
	r2 := partTwo(lines)
	fmt.Println(r2)
}

func fillRucksacks(lines []string) []Rucksack {
	var rucksacks []Rucksack
	for _, line := range lines {
		compartment_size := len(line) / 2
		left := line[:compartment_size]
		right := line[compartment_size:]
		rucksacks = append(rucksacks, Rucksack{left, right})
	}

	return rucksacks
}

func partOne(rucksacks []Rucksack) int {
	total := 0
	for _, rucksack := range rucksacks {
		common := getCommon(rucksack)
		total += getItemValue(common)
	}
	return total
}

func getCommon(rucksack Rucksack) rune {
	for _, c := range rucksack.left {
		if strings.Contains(rucksack.right, string(c)) {
			return c
		}
	}
	return -1
}

func getItemValue(item rune) int {
	// a-z = 97-122
	// A-Z = 65-90
	i := int(item)
	if i >= 65 && i <= 90 {
		return i - 64 + 26
	}

	return i - 96
}

func partTwo(lines []string) int {
	// brute force
	total := 0
	for i, _ := range lines {
		if i%3 == 0 {
			common := getCommonBetweenRucksacks(lines[i], lines[i+1], lines[i+2])
			total += getItemValue(common)
		}
	}

	return total
}

func getCommonBetweenRucksacks(r1, r2, r3 string) rune {
	for _, c := range r1 {
		if strings.Contains(r2, string(c)) && strings.Contains(r3, string(c)) {
			return c
		}
	}
	return -1
}
