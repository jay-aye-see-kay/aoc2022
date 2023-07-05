package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("part 1: %+v\n", part1())
	fmt.Printf("part 2: %+v\n", part2())
}

func readInput() string {
	content, err := ioutil.ReadFile("./inputs/day08.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	return strings.TrimSpace(string(content))
}

type Grid [][]int

func (grid Grid) maxX() int {
	return len(grid[0]) - 1
}
func (grid Grid) maxY() int {
	return len(grid[0]) - 1
}

// 3851: wrong
// 3981: wrong
func part1() int {
	grid := parse(readInput())
	count := 0

	for y := 0; y <= grid.maxY(); y++ {
		for x := 0; x <= grid.maxX(); x++ {
			if grid.isVisible(x, y) {
				count++
			}
		}
	}

	return count
}

func part2() int {
	return 0
}

func parse(input string) Grid {
	grid := Grid{}
	for lineIdx, line := range strings.Split(input, "\n") {
		grid = append(grid, []int{})
		for _, char := range strings.Split(line, "") {
			height, _ := strconv.Atoi(char)
			grid[lineIdx] = append(grid[lineIdx], height)
		}
	}
	return grid
}

func (grid Grid) isVisible(x, y int) bool {
	height := grid[y][x]

	// if edge return true
	if x == 0 || x == grid.maxX() || y == 0 || y == grid.maxY() {
		return true
	}

	// loop and check in all 4 directions

	// loop to the left
	leftVisible := true
	for i := x; i > 0; i-- {
		if grid[y][i] > height {
			leftVisible = false
		}
	}
	if leftVisible {
		return true
	}

	// loop to the right
	rightVisible := true
	for i := x; i < grid.maxX(); i++ {
		if grid[y][i] > height {
			rightVisible = false
		}
	}
	if rightVisible {
		return true
	}

	// loop to the bottom
	bottomVisible := true
	for i := y; i < grid.maxY(); i++ {
		if grid[i][x] > height {
			bottomVisible = false
		}
	}
	if bottomVisible {
		return true
	}

	// loop to the top
	topVisible := true
	for i := y; i > 0; i-- {
		if grid[i][x] > height {
			topVisible = false
		}
	}
	if topVisible {
		return true
	}

	return false
}
