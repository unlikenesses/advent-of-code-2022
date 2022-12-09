package main

import (
	"fmt"
	"strings"

	"github.com/unlikenesses/utils"
)

func main() {
	lines := utils.ReadInput()
	r1 := partOne(lines[0], 4)
	fmt.Println(r1)
	r2 := partOne(lines[0], 14)
	fmt.Println(r2)
}

func partOne(stream string, numChars int) int {
	for i := 0; i < len(stream)-(numChars-1); i++ {
		if allDifferent(stream[i : i+numChars]) {
			return i + numChars
		}
	}
	return 0
}

func allDifferent(chars string) bool {
	for _, c := range chars {
		if strings.Count(chars, string(c)) > 1 {
			return false
		}
	}
	return true
}
