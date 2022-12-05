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
	content, err := ioutil.ReadFile("../inputs/day05.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	return string(content)
}

type StrList []string

func (s *StrList) pop() string {
	sl := *s
	last := sl[len(sl)-1]
	*s = sl[:len(sl)-1]
	return last
}
func (s *StrList) push(el string) {
	*s = append(*s, el)
}
func (s *StrList) len() int {
	return len(*s)
}
func makeStrList(s string) StrList {
	list := strings.Split(s, "")
	// reverse it (hack, remove when writing proper parser)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}

func part1() string {
	stacks, moves := parse()
	for _, move := range moves {
		for i := 0; i < move.count; i++ {
			stacks[move.to].push(stacks[move.from].pop())
		}
	}
	answer := ""
	for _, stack := range stacks {
		answer += stack.pop()
	}
	return answer
}

func part2() string {
	stacks, moves := parse()
	for _, move := range moves {
		toPush := StrList{}
		for i := 0; i < move.count; i++ {
			toPush.push(stacks[move.from].pop())
		}
		for toPush.len() > 0 {
			stacks[move.to].push(toPush.pop())
		}
	}
	answer := ""
	for _, stack := range stacks {
		answer += stack.pop()
	}
	return answer
}

type Move struct {
	count int
	from  int
	to    int
}

func parse() ([]StrList, []Move) {
	stacks := []StrList{
		makeStrList("SPHVFG"),
		makeStrList("MZDVBFJG"),
		makeStrList("NJLMG"),
		makeStrList("PWDVZGN"),
		makeStrList("BCRV"),
		makeStrList("ZLWPMSRV"),
		makeStrList("PHT"),
		makeStrList("VZHCNSRQ"),
		makeStrList("JQVPGLF"),
	}
	input := readInput()
	halves := strings.Split(input, "\n\n")
	movesStr := strings.TrimSpace(halves[1])
	moves := []Move{}
	for _, line := range strings.Split(movesStr, "\n") {
		bits := strings.Split(line, " ")
		count, _ := strconv.Atoi(bits[1])
		from, _ := strconv.Atoi(bits[3])
		to, _ := strconv.Atoi(bits[5])
		moves = append(moves, Move{count, from - 1, to - 1})
	}
	return stacks, moves
}
