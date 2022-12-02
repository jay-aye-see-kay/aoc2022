package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("day 1 part 1: %+v\n", part1())
	fmt.Printf("day 1 part 2: %+v\n", part2())
}

func readInput() string {
	content, err := ioutil.ReadFile("../inputs/day02.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	return strings.TrimSpace(string(content))
}

// A for Rock, B for Paper, and C for Scissors
// X for Rock, Y for Paper, and Z for Scissors

// points:
// 1 for Rock, 2 for Paper, and 3 for Scissors (shape you selected)
// 0 if you lost, 3 if the round was a draw, and 6 if you won

func part1() int {
	text := readInput()
	sum := 0
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)

		result := map[string]string{
			"A X": "draw",
			"A Y": "win",
			"A Z": "lose",
			"B X": "lose",
			"B Y": "draw",
			"B Z": "win",
			"C X": "win",
			"C Y": "lose",
			"C Z": "draw",
		}[line]

		resultPoints := map[string]int{
			"lose": 0,
			"draw": 3,
			"win":  6,
		}[result]

		choice := string(line[2])
		choicePoints := map[string]int{
			"X": 1,
			"Y": 2,
			"Z": 3,
		}[choice]

		sum += resultPoints + choicePoints
	}

	return sum
}

// X means you need to lose, Y means you need to end the round in a draw,
// and Z means you need to win
func part2() int {
	text := readInput()
	sum := 0
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)

		myChoice := map[string]string{
			"A X": "scissors",
			"A Y": "rock",
			"A Z": "paper",
			"B X": "rock",
			"B Y": "paper",
			"B Z": "scissors",
			"C X": "paper",
			"C Y": "scissors",
			"C Z": "rock",
		}[line]

    choicePoints := map[string]int{
      "rock": 1,
      "paper": 2,
      "scissors": 3,
    }[myChoice]

		requiredResult := string(line[2])
		resultPoints := map[string]int{
			"X": 0,
			"Y": 3,
			"Z": 6,
		}[requiredResult]

		sum += resultPoints + choicePoints
	}

	return sum
}
