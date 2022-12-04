package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("part 1: %+v\n", part1())
	// fmt.Printf("part 2: %+v\n", part2())
}

func readInput() string {
	content, err := ioutil.ReadFile("../inputs/day04.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	return strings.TrimSpace(string(content))
}

func part1() int {
	assignments := parseInput(readInput())
	count := 0
	for _, ass := range assignments {
		if completelyOverlaps(ass) {
			count++
		}
	}
	return count
}

func part2() int {
	assignments := parseInput(readInput())
	count := 0
	for _, ass := range assignments {
		if partiallyOverlaps(ass) {
			count++
		}
	}
	return count
}

func completelyOverlaps(ass []int) bool {
	xStart := ass[0]
	xEnd := ass[1]
	yStart := ass[2]
	yEnd := ass[3]

	if xStart >= yStart && xEnd <= yEnd {
		return true
	} else if yStart >= xStart && yEnd <= xEnd {
		return true
	}
	return false
}

func partiallyOverlaps(ass []int) bool {
	xStart := ass[0]
	xEnd := ass[1]
	yStart := ass[2]
	yEnd := ass[3]

	if xEnd < yStart || xStart > yEnd {
		return false
	}
	return true
}

func parseInput(s string) [][]int {
	lines := strings.Split(s, "\n")
	out := make([][]int, len(lines))

	for i, line := range lines {
		halves := strings.Split(line, ",")
		left := halves[0]
		leftQuarter := strings.Split(left, "-")
		leftStart, _ := strconv.Atoi(leftQuarter[0])
		leftEnd, _ := strconv.Atoi(leftQuarter[1])

		right := halves[1]
		rightQuarter := strings.Split(right, "-")
		rightStart, _ := strconv.Atoi(rightQuarter[0])
		rightEnd, _ := strconv.Atoi(rightQuarter[1])

		out[i] = []int{leftStart, leftEnd, rightStart, rightEnd}
	}

	return out
}
