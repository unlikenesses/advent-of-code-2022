package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/unlikenesses/utils"
)

type Instruction struct {
	qty  int
	src  int
	dest int
}

var stacks = make(map[int][]string)
var numStacks int
var instructions []Instruction

func main() {
	lines := utils.ReadInput()
	stacks, instructions = parseStacks(lines)
	r1 := partOne(true)
	fmt.Println(r1)
	// clean map
	for i := 1; i <= numStacks; i++ {
		stacks[i] = []string{}
	}
	stacks, instructions = parseStacks(lines)
	r2 := partOne(false)
	fmt.Println(r2)
}

func parseStacks(lines []string) (map[int][]string, []Instruction) {
	var instructions []Instruction
	inStacks := true
	numStacks = getNumStacks(lines)
	re := regexp.MustCompile(`move (\d+) from (\d) to (\d)`)
	for _, line := range lines {
		if len(line) > 0 && line[1] == '1' {
			inStacks = false
			continue
		}
		if inStacks {
			for i := 1; i <= numStacks; i++ {
				crate := readCrate(line, i)
				if crate != " " {
					stacks[i] = append(stacks[i], crate)
				}
			}
		} else {
			// read instructions
			matches := re.FindStringSubmatch(line)
			if len(matches) > 0 {
				qty, _ := strconv.Atoi(matches[1])
				src, _ := strconv.Atoi(matches[2])
				dest, _ := strconv.Atoi(matches[3])
				instruction := Instruction{qty, src, dest}
				instructions = append(instructions, instruction)
			} else {
				fmt.Println("No matches for line ", line)
			}
		}
	}

	return stacks, instructions
}

func getNumStacks(lines []string) int {
	for _, line := range lines {
		if line[1] == '1' {
			lastStack := string(line[len(line)-2])
			lastStackInt, _ := strconv.Atoi(lastStack)
			return lastStackInt
		}
	}
	return 0
}

func readCrate(line string, pos int) string {
	cratePos := ((pos - 1) * 4) + 1
	return string(line[cratePos])
}

func partOne(oldCrane bool) string {
	var message string
	for _, instruction := range instructions {
		srcStack := stacks[instruction.src]
		destStack := stacks[instruction.dest]
		toMove := srcStack[:instruction.qty]
		if oldCrane {
			toMove = utils.ReverseSlice(toMove)
		} else {
			// yuck
			toMove = append([]string{}, toMove...)
		}
		stacks[instruction.dest] = append(toMove, destStack...)
		var newSrc []string
		for i := 0; i < len(srcStack); i++ {
			if i >= instruction.qty {
				newSrc = append(newSrc, srcStack[i])
			}
		}
		stacks[instruction.src] = newSrc
	}
	for i := 1; i <= numStacks; i++ {
		message += stacks[i][0]
	}
	return message
}
