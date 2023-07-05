import * as fs from "fs/promises";
import { part1, part2 } from "./04";

describe("day 4", async () => {
  const realInput = await fs.readFile("../inputs/day04.txt", "utf8");

  test("part 1 test input", () => {
    expect(part1(testInput)).toBe(2);
  });

  test("part 1 real input", () => {
    expect(part1(realInput)).toBe(433);
  });

  test.skip("part 2 test input", () => {
    expect(part2(testInput)).toBe(0);
  });

  test.skip("part 2 real input", () => {
    expect(part2(realInput)).toBe(0);
  });
});

const testInput = `
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`;
