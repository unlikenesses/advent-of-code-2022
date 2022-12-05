package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/unlikenesses/utils"
)

type Pair struct {
	first  []int
	second []int
}

func main() {
	lines := utils.ReadInput()
	pairs := parsePairs(lines)
	r1 := partOne(pairs)
	fmt.Println(r1)
	r2 := partTwo(pairs)
	fmt.Println(r2)
}

func parsePairs(lines []string) []Pair {
	var pairs []Pair
	for _, line := range lines {
		split := strings.Split(line, ",")
		pairs = append(pairs, Pair{parseRange(split[0]), parseRange(split[1])})
	}

	return pairs
}

func parseRange(input string) []int {
	var output []int
	split := strings.Split(input, "-")
	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])
	for i := start; i <= end; i++ {
		output = append(output, i)
	}
	return output
}

func partOne(pairs []Pair) int {
	total := 0
	for _, pair := range pairs {
		if oneContainsOther(pair) {
			total++
		}
	}
	return total
}

func oneContainsOther(pair Pair) bool {
	// if both are same length, nope:
	if len(pair.first) == len(pair.second) && (!reflect.DeepEqual(pair.first, pair.second)) {
		return false
	}
	if len(pair.first) > len(pair.second) {
		return contains(pair.first, pair.second)
	}
	return contains(pair.second, pair.first)
}

func contains(larger, smaller []int) bool {
	smallerStart := smaller[0]
	smallerEnd := smaller[len(smaller)-1]
	largerStart := larger[0]
	largerEnd := larger[len(larger)-1]
	return smallerStart >= largerStart && smallerEnd <= largerEnd
}

func partTwo(pairs []Pair) int {
	total := 0
	for _, pair := range pairs {
		if overlaps(pair) {
			total++
		}
	}
	return total
}

func overlaps(pair Pair) bool {
	firstStart := pair.first[0]
	firstEnd := pair.first[len(pair.first)-1]
	secondStart := pair.second[0]
	secondEnd := pair.second[len(pair.second)-1]
	return firstEnd >= secondStart && firstStart <= secondEnd
}
