package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/unlikenesses/utils"
)

type Node struct {
	name     string
	parent   *Node
	isDir    bool
	size     int
	children []*Node
}

var sizes []int

func main() {
	lines := utils.ReadInput()
	n := Node{"/", &Node{}, true, 0, []*Node{}}
	n = buildTree(lines[1:], n)
	// printTree(&n, 0)
	r1 := partOne(&n)
	fmt.Println(r1)
	r2 := partTwo(&n)
	fmt.Println(r2)
}

func buildTree(lines []string, n Node) Node {
	var name string
	var size int
	var isDir bool
	var currentDir *Node
	currentDir = &n
	for _, line := range lines {
		// cd
		re := regexp.MustCompile(`\$ cd (\/|\.\.|[a-z]+)*`)
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			dir := matches[1]
			if dir == ".." {
				// move currentDir up one
				currentDir = currentDir.parent
			} else {
				// have to find this directory in current node's children
				for _, cn := range currentDir.children {
					if cn.name == dir {
						currentDir = cn
						break
					}
				}
			}
			continue
		}
		// ls - we can ignore this
		if line == "$ ls" {
			continue
		}
		// an element in the directory listing
		re = regexp.MustCompile(`dir ([a-z]+)`)
		matches = re.FindStringSubmatch(line)
		if len(matches) > 0 {
			isDir = true
			size = 0
			name = matches[1]
		} else {
			re = regexp.MustCompile(`(\d+) ([a-z|.]+)`)
			matches = re.FindStringSubmatch(line)
			if len(matches) > 0 {
				isDir = false
				size, _ = strconv.Atoi(matches[1])
				name = matches[2]
			} else {
				panic("Something went wrong parsing the files")
			}
		}
		childNode := Node{name, currentDir, isDir, size, []*Node{}}
		currentDir.children = append(currentDir.children, &childNode)
	}
	return n
}

func printTree(n *Node, d int) {
	for i := 0; i < d; i++ {
		fmt.Print("|")
	}
	name := n.name
	if n.isDir {
		name += " (dir)"
	}
	fmt.Println(name)
	for _, c := range n.children {
		printTree(c, d+1)
	}
}

func partOne(n *Node) int {
	getSizes(n, true)
	return utils.GetSum(sizes)
}

func partTwo(n *Node) int {
	getSizes(n, false)
	sort.Ints(sizes)
	diskSpace := 70000000
	target := 30000000
	usedSpace := sizes[len(sizes)-1]
	unusedSpace := diskSpace - usedSpace
	spaceToFree := target - unusedSpace
	for _, size := range sizes {
		if size >= spaceToFree {
			return size
		}
	}
	return 0
}

func getSizes(n *Node, limit bool) int {
	var size int
	for _, c := range n.children {
		if c.isDir {
			size += getSizes(c, limit)
			continue
		}
		size += c.size
	}
	if !limit || size <= 100000 {
		sizes = append(sizes, size)
	}
	return size
}
