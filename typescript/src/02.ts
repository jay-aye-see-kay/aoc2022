import * as R from "ramda";

function play(input: string, scoreTable: Record<string, number>) {
  return input
    .trim()
    .split("\n")
    .map((line) => scoreTable[line])
    .reduce((sum, n) => sum + n, 0);
}

// The first column is what your opponent is going to play:
// A for Rock, B for Paper, and C for Scissors.

// The second column is what you should play in response:
// X for Rock, Y for Paper, and Z for Scissors

// The score for a single round is the score for the shape you selected
// (1 for Rock, 2 for Paper, and 3 for Scissors)

// plus the score for the outcome of the round
// (0 if you lost, 3 if the round was a draw, and 6 if you won).

export function part1(input: string) {
  const scoreTable = {
    "A X": 3 + 1,
    "A Y": 6 + 2,
    "A Z": 0 + 3,
    "B X": 0 + 1,
    "B Y": 3 + 2,
    "B Z": 6 + 3,
    "C X": 6 + 1,
    "C Y": 0 + 2,
    "C Z": 3 + 3,
  };
  return play(input, scoreTable);
}

// the second column says how the round needs to end:
// X means you need to lose
// Y means you need to end the round in a draw
// Z means you need to win

export function part2(input: string) {
  const scoreTable = {
    "A X": 0 + 3,
    "A Y": 3 + 1,
    "A Z": 6 + 2,
    "B X": 0 + 1,
    "B Y": 3 + 2,
    "B Z": 6 + 3,
    "C X": 0 + 2,
    "C Y": 3 + 3,
    "C Z": 6 + 1,
  };
  return play(input, scoreTable);
}
