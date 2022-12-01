package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("day 1 part 1: %+v\n", Part1())
	fmt.Printf("day 1 part 2: %+v\n", Part2())
}

func readInput() string {
	content, err := ioutil.ReadFile("inputs/day01.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	return strings.TrimSpace(string(content))
}

func parseInput(text string) []int {
	groups := strings.Split(text, "\n\n")
	elfTotals := []int{}

	for _, group := range groups {
		lines := strings.Split(group, "\n")
		sum := 0
		for _, line := range lines {
			num, _ := strconv.Atoi(line)
			sum += num
		}
		elfTotals = append(elfTotals, sum)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfTotals)))
	return elfTotals
}

func Part1() int {
	input := readInput()
	elfTotals := parseInput(input)
	return elfTotals[0]
}

func Part2() int {
	input := readInput()
	elfTotals := parseInput(input)
	return elfTotals[0] + elfTotals[1] + elfTotals[2]
}
