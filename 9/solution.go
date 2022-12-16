package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/unlikenesses/utils"
)

type Coord struct {
	x, y int
}

var lines []string

func main() {
	lines = utils.ReadInput()
	r1 := partOne(1)
	fmt.Println(r1)
	r2 := partOne(9)
	fmt.Println(r2)
}

func partOne(numTailSegments int) int {
	var headPos, lead Coord
	tailSegments := make([]Coord, numTailSegments)
	var tailPositions []Coord
	re := regexp.MustCompile(`(L|R|U|D) (\d+)`)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			dir := matches[1]
			qty, _ := strconv.Atoi(matches[2])
			if dir == "R" {
				for i := 0; i < qty; i++ {
					headPos.x++
					for t, tailPos := range tailSegments {
						if t == 0 {
							lead = headPos
						} else {
							lead = tailSegments[t-1]
						}
						tx, ty := getTailPos(lead, tailPos, dir)
						tailPos.x = tx
						tailPos.y = ty
						tailSegments[t] = tailPos
						if t == len(tailSegments)-1 && !posInArray(tailPos, tailPositions) {
							tailPositions = append(tailPositions, tailPos)
						}
					}
				}
			} else if dir == "L" {
				for i := 0; i < qty; i++ {
					headPos.x--
					for t, tailPos := range tailSegments {
						if t == 0 {
							lead = headPos
						} else {
							lead = tailSegments[t-1]
						}
						tx, ty := getTailPos(lead, tailPos, dir)
						tailPos.x = tx
						tailPos.y = ty
						tailSegments[t] = tailPos
						if t == len(tailSegments)-1 && !posInArray(tailPos, tailPositions) {
							tailPositions = append(tailPositions, tailPos)
						}
					}
				}
			} else if dir == "U" {
				for i := 0; i < qty; i++ {
					headPos.y++
					for t, tailPos := range tailSegments {
						if t == 0 {
							lead = headPos
						} else {
							lead = tailSegments[t-1]
						}
						tx, ty := getTailPos(lead, tailPos, dir)
						tailPos.x = tx
						tailPos.y = ty
						tailSegments[t] = tailPos
						if t == len(tailSegments)-1 && !posInArray(tailPos, tailPositions) {
							tailPositions = append(tailPositions, tailPos)
						}
					}
				}
			} else if dir == "D" {
				for i := 0; i < qty; i++ {
					headPos.y--
					for t, tailPos := range tailSegments {
						if t == 0 {
							lead = headPos
						} else {
							lead = tailSegments[t-1]
						}
						tx, ty := getTailPos(lead, tailPos, dir)
						tailPos.x = tx
						tailPos.y = ty
						tailSegments[t] = tailPos
						if t == len(tailSegments)-1 && !posInArray(tailPos, tailPositions) {
							tailPositions = append(tailPositions, tailPos)
						}
					}
				}
			}
		} else {
			fmt.Println("No matches for line ", line)
		}
	}
	return len(tailPositions)
}

func getTailPos(headPos, tailPos Coord, dir string) (x, y int) {
	if touching(headPos, tailPos) {
		return tailPos.x, tailPos.y
	}
	// on same row with a gap between them
	if headPos.y == tailPos.y {
		if headPos.x > tailPos.x {
			return tailPos.x + 1, tailPos.y
		}
		return tailPos.x - 1, tailPos.y
	}
	// on same column with a gap between them
	if headPos.x == tailPos.x {
		if headPos.y > tailPos.y {
			return tailPos.x, tailPos.y + 1
		}
		return tailPos.x, tailPos.y - 1
	}
	// only option left is diagonals
	if headPos.y > tailPos.y {
		if headPos.x > tailPos.x {
			return tailPos.x + 1, tailPos.y + 1
		}
		return tailPos.x - 1, tailPos.y + 1
	}

	if headPos.y < tailPos.y {
		if headPos.x > tailPos.x {
			return tailPos.x + 1, tailPos.y - 1
		}
		return tailPos.x - 1, tailPos.y - 1
	}
	return 0, 0
}

func touching(headPos, tailPos Coord) bool {
	// on same spot
	if headPos.x == tailPos.x && headPos.y == tailPos.y {
		return true
	}
	// 1 to the left or right
	if headPos.y == tailPos.y && utils.IntAbs(headPos.x-tailPos.x) < 2 {
		return true
	}
	// 1 above or below
	if headPos.x == tailPos.x && utils.IntAbs(headPos.y-tailPos.y) < 2 {
		return true
	}
	// diagonally adjacent
	if utils.IntAbs(headPos.x-tailPos.x) < 2 && utils.IntAbs(headPos.y-tailPos.y) < 2 {
		return true
	}
	return false
}

func posInArray(needle Coord, haystack []Coord) bool {
	for _, i := range haystack {
		if i == needle {
			return true
		}
	}
	return false
}
