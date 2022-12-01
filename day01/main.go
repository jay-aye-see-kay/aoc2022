package day01

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func readInput() string {
	content, err := ioutil.ReadFile("inputs/day01.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	return strings.TrimSpace(string(content))
}

type Elf = []int

func parseInput(text string) []Elf {
	groups := strings.Split(text, "\n\n")
	elves := [][]int{}

	for groupId, group := range groups {
		lines := strings.Split(group, "\n")
		elves = append(elves, []int{})

		for _, line := range lines {
			num, _ := strconv.Atoi(line)
			elves[groupId] = append(elves[groupId], num)
		}
	}
	return elves
}

func Part1() int {
	input := readInput()
	elves := parseInput(input)

	max := 0
	for _, elf := range elves {
		elfSum := 0
		for _, cal := range elf {
			elfSum += cal
		}
		if elfSum > max {
			max = elfSum
		}
	}
	return max
}

func Part2() int {
	input := readInput()
	elves := parseInput(input)

	elfCals := []int{}
	for _, elf := range elves {
		elfSum := 0
		for _, cal := range elf {
			elfSum += cal
		}
		elfCals = append(elfCals, elfSum)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfCals)))

	return elfCals[0] + elfCals[1] + elfCals[2]
}
