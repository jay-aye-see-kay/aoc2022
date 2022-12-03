package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("part 1: %+v\n", part1())
	fmt.Printf("part 2: %+v\n", part2())
}

func readInput() string {
	content, err := ioutil.ReadFile("../inputs/day03.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	return strings.TrimSpace(string(content))
}

func part1() int {
	sum := 0
	for _, line := range strings.Split(readInput(), "\n") {
		sum += getItemValue(findCommon2(line))
	}
	return sum
}

func part2() int {
	lines := strings.Split(readInput(), "\n")
	sum := 0
	for i := 0; i < len(lines)-2; i += 3 {
		sum += getItemValue(findCommon3(lines[i], lines[i+1], lines[i+2]))
	}
	return sum
}

func getItemValue(item string) int {
	value := int(item[0])
	if value >= 65 && value <= 90 { // capital letter
		return value - 38
	} else if value >= 97 && value <= 122 { // lower case letter
		return value - 96
	}
	panic("unexpected letter")
}

func findCommon2(line string) string {
	mid := len(line) / 2
	left := line[:mid]
	right := line[mid:]
	for _, char := range strings.Split(left, "") {
		if strings.Contains(right, char) {
			return char
		}
	}
	panic("could not find common letter")
}

func findCommon3(line1, line2, line3 string) string {
	for _, char := range strings.Split(line1, "") {
		if strings.Contains(line2, char) && strings.Contains(line3, char) {
			return char
		}
	}
	panic("could not find common letter")
}
