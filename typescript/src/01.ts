import * as R from "ramda";

export function part1(input: string) {
  const groups = input
    .trim()
    .split("\n\n")
    .map((group) => group.split("\n").map((line) => parseInt(line)));

  const sums = groups.map(R.sum);
  const max = Math.max(...sums);
  return max;
}
