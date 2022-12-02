package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/unlikenesses/utils"
)

func main() {
	lines := utils.ReadInput()
	totals := getTotals(lines)
	r1 := partOne(totals)
	fmt.Println(r1)
	r2 := partTwo(totals)
	fmt.Println(r2)
}

func getTotals(lines []string) []int {
	var totals []int
	total := 0
	for _, line := range lines {
		if line != "" {
			cals, _ := strconv.Atoi(line)
			total += cals
		} else {
			totals = append(totals, total)
			total = 0
		}
	}
	totals = append(totals, total)

	return totals
}

func partOne(totals []int) int {
	return utils.GetMax(totals)
}

func partTwo(totals []int) int {
	sort.Ints(totals)

	return utils.GetSum(totals[len(totals)-3:])
}
