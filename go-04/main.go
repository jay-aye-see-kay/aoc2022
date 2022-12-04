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
	content, err := ioutil.ReadFile("../inputs/day04.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	return strings.TrimSpace(string(content))
}

func part1() int {
	count := 0
	for _, assignment := range parseInput(readInput()) {
		if assignment.completelyOverlap() {
			count++
		}
	}
	return count
}

func part2() int {
	count := 0
	for _, assignment := range parseInput(readInput()) {
		if assignment.partiallyOverlap() {
			count++
		}
	}
	return count
}

type TimeRange struct {
	start int
	end   int
}

type Assignment struct {
	x TimeRange
	y TimeRange
}

func (a *Assignment) completelyOverlap() bool {
	return (a.x.start >= a.y.start && a.x.end <= a.y.end) ||
		(a.y.start >= a.x.start && a.y.end <= a.x.end)
}

func (a *Assignment) partiallyOverlap() bool {
	return !(a.x.end < a.y.start || a.x.start > a.y.end)
}

func parseInput(s string) []Assignment {
	lines := strings.Split(s, "\n")
	out := make([]Assignment, len(lines))

	for i, line := range lines {
		nums := make([]int, 4)
		for i, numStr := range strings.Split(strings.ReplaceAll(line, ",", "-"), "-") {
			nums[i], _ = strconv.Atoi(numStr)
		}
		out[i] = Assignment{TimeRange{nums[0], nums[1]}, TimeRange{nums[2], nums[3]}}
	}

	return out
}
