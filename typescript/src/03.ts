import * as R from "ramda";

export function part1(input: string) {
  const rucksacks = input
    .trim()
    .split("\n")
    .map((line) => [
      line.slice(0, line.length / 2),
      line.slice(line.length / 2),
    ]);

  const values = rucksacks.map(findCommonChar).map(getCharValue);
  return R.sum(values);
}

export function part2(input: string) {
  // group lines into 3s to form rucksacks
  const rucksacks: string[][] = [];
  for (const [i, line] of input.trim().split("\n").entries()) {
    if (i % 3 == 0) rucksacks.push([]);
    rucksacks.at(-1)?.push(line);
  }

  const values = rucksacks.map(findCommonChar).map(getCharValue);
  return R.sum(values);
}

/** Finds the common character in a list of strings */
function findCommonChar(strings: string[]) {
  let charGroups = strings.map((s) => s.split(""));
  while (charGroups.length >= 2) {
    const [s1, s2, ...rest] = charGroups;
    charGroups = [R.intersection(s1, s2), ...rest];
  }
  return charGroups[0][0];
}

function getCharValue(item: string) {
  const charCode = item.charCodeAt(0);
  if (charCode >= 65 && charCode <= 90) {
    return charCode - 38;
  } else if (charCode >= 97 && charCode <= 122) {
    return charCode - 96;
  }
  return 0;
}
