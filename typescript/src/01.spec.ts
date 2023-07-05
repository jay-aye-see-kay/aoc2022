import * as fs from "fs/promises";
import { part1 } from "./01";

test("part 1 test input", () => {
  expect(part1(testInput)).toBe(24000);
});

test("part 1 real input", async () => {
  const realInput = await fs.readFile("../inputs/day01.txt", "utf8");
  expect(part1(realInput)).toBe(71934);
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
