import * as R from "ramda";

export function part1(input: string) {
  const rucksacks = input
    .trim()
    .split("\n")
    .map((line) => [
      line.slice(0, line.length / 2),
      line.slice(line.length / 2),
    ]);

  const values = rucksacks
    .map(([c1, c2]) => R.intersection(c1.split(""), c2.split(""))[0])
    .map(getValue);
  return R.sum(values);
}

export function part2(input: string) {
  return 0;
}

function getValue(item: string) {
  const charCode = item.charCodeAt(0);
  if (charCode >= 65 && charCode <= 90) {
    return charCode - 38;
  } else if (charCode >= 97 && charCode <= 122) {
    return charCode - 96;
  }
  return 0;
}
