import * as fs from "fs/promises";
import { part1, part2 } from "./03";

describe("day 3", async () => {
  const realInput = await fs.readFile("../inputs/day03.txt", "utf8");

  test("part 1 test input", () => {
    expect(part1(testInput)).toBe(157);
  });

  test("part 1 real input", () => {
    expect(part1(realInput)).toBe(7428);
  });

  test.skip("part 2 test input", () => {
    expect(part2(testInput)).toBe(70);
  });

  test.skip("part 2 real input", () => {
    expect(part2(realInput)).toBe(2650);
  });
});

const testInput = `
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`;
