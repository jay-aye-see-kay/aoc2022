import * as R from "ramda";

function sumGroups(input: string) {
  const sums = input
    .trim()
    .split("\n\n")
    .map((group) => group.split("\n").map((line) => parseInt(line)))
    .map(R.sum);

  return R.sort((a, b) => b - a, sums);
}

export function part1(input: string) {
  return sumGroups(input)[0];
}

export function part2(input: string) {
  const groups = sumGroups(input);
  return groups[0] + groups[1] + groups[2];
}
