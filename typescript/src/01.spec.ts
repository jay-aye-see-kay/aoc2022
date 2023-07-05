import * as fs from "fs/promises";
import { part1, part2 } from "./01";

describe("day 1", async () => {
  const realInput = await fs.readFile("../inputs/day01.txt", "utf8");

  test("part 1 test input", () => {
    expect(part1(testInput)).toBe(24000);
  });

  test("part 1 real input", () => {
    expect(part1(realInput)).toBe(71934);
  });

  test("part 2 test input", () => {
    expect(part2(testInput)).toBe(45000);
  });

  test("part 2 real input", () => {
    expect(part2(realInput)).toBe(211447);
  });
});

const testInput = `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`;
