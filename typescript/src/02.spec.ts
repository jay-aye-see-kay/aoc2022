import * as fs from "fs/promises";
import { part1, part2 } from "./02";

describe("day 2", async () => {
  const realInput = await fs.readFile("../inputs/day02.txt", "utf8");

  test("part 1 test input", () => {
    expect(part1(testInput)).toBe(15);
  });

  test("part 1 real input", () => {
    expect(part1(realInput)).toBe(11906);
  });

  test("part 2 test input", () => {
    expect(part2(testInput)).toBe(12);
  });

  test("part 2 real input", () => {
    expect(part2(realInput)).toBe(11186);
  });
});

const testInput = `
A Y
B X
C Z
`;
