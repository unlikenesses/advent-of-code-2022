package main

import (
	"fmt"
	"strconv"

	"github.com/unlikenesses/utils"
)

var lines []string

func main() {
	lines = utils.ReadInput()
	r1 := partOne()
	fmt.Println(r1)
	r2 := partTwo()
	fmt.Println(r2)
}

func partOne() int {
	var numVisible int
	for y, line := range lines {
		if y == 0 || y == len(lines)-1 {
			numVisible += len(lines)
			continue
		}
		for x, t := range line {
			if x == 0 || x == len(line)-1 {
				numVisible++
				continue
			}
			treeHeight, _ := strconv.Atoi(string(t))
			if isVisible(x, y, treeHeight) {
				numVisible++
			}
		}
	}
	return numVisible
}

func isVisible(x, y, height int) bool {
	// north
	northVisible := true
	for i := y - 1; i >= 0; i-- {
		if getTreeHeight(x, i) >= height {
			northVisible = false
		}
	}
	// south
	southVisible := true
	for i := y + 1; i < len(lines); i++ {
		if getTreeHeight(x, i) >= height {
			southVisible = false
		}
	}
	// east
	eastVisible := true
	for i := x + 1; i < len(lines[0]); i++ {
		if getTreeHeight(i, y) >= height {
			eastVisible = false
		}
	}
	// west
	westVisible := true
	for i := x - 1; i >= 0; i-- {
		if getTreeHeight(i, y) >= height {
			westVisible = false
		}
	}
	return (northVisible || southVisible || eastVisible || westVisible)
}

func partTwo() int {
	var bestScore int
	for y, line := range lines {
		for x, t := range line {
			treeHeight, _ := strconv.Atoi(string(t))
			score := getScore(x, y, treeHeight)
			if score > bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
}

func getScore(x, y, height int) int {
	// north
	var northScore int
	if y > 0 {
		for i := y - 1; i >= 0; i-- {
			h := getTreeHeight(x, i)
			if h < height || (h == height && h == 0) {
				northScore++
				continue
			}
			if h >= height {
				northScore++
				break
			}
		}
	}
	// south
	var southScore int
	if y < len(lines)-1 {
		for i := y + 1; i < len(lines); i++ {
			h := getTreeHeight(x, i)
			if h < height || (h == height && h == 0) {
				southScore++
				continue
			}
			if h >= height {
				southScore++
				break
			}
		}
	}
	// east
	var eastScore int
	if x < len(lines[0])-1 {
		for i := x + 1; i < len(lines[0]); i++ {
			h := getTreeHeight(i, y)
			if h < height || (h == height && h == 0) {
				eastScore++
				continue
			}
			if h >= height {
				eastScore++
				break
			}
		}
	}
	// west
	var westScore int
	if x > 0 {
		for i := x - 1; i >= 0; i-- {
			h := getTreeHeight(i, y)
			if h < height || (h == height && h == 0) {
				westScore++
				continue
			}
			if h >= height {
				westScore++
				break
			}
		}
	}
	return northScore * southScore * eastScore * westScore
}

func getTreeHeight(x, y int) int {
	c := lines[y][x]
	height, _ := strconv.Atoi(string(c))
	return height
}
