package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("day 1 part 1: %+v\n", part1())
	fmt.Printf("day 1 part 2: %+v\n", part2())
}

func readInput() string {
	content, err := ioutil.ReadFile("../inputs/day01.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	return strings.TrimSpace(string(content))
}

func parseInput(text string) []int {
	groups := strings.Split(text, "\n\n")
	elfTotals := make([]int, len(groups))

	for i, group := range groups {
		lines := strings.Split(group, "\n")
		for _, line := range lines {
			num, _ := strconv.Atoi(line)
			elfTotals[i] += num
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfTotals)))
	return elfTotals
}

func part1() int {
	input := readInput()
	elfTotals := parseInput(input)
	return elfTotals[0]
}

func part2() int {
	input := readInput()
	elfTotals := parseInput(input)
	return elfTotals[0] + elfTotals[1] + elfTotals[2]
}
