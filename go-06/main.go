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
	content, err := ioutil.ReadFile("../inputs/day06.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	return strings.TrimSpace(string(content))
}

type RuneSet map[rune]bool

func makeRuneSet(str string) RuneSet {
	set := RuneSet{}
	for _, char := range str {
		set[char] = true
	}
	return set
}

func findUniqSubstr(input_str string, substr_len int) int {
	for i := substr_len; i < len(input_str); i++ {
		set := makeRuneSet(input_str[i-substr_len : i])
		if len(set) == substr_len {
			return i
		}
	}
	panic("not found")
}

func part1() int {
	return findUniqSubstr(readInput(), 4)
}

func part2() int {
	return findUniqSubstr(readInput(), 14)
}
