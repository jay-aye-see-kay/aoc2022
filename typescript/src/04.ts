import * as R from "ramda";

export function part1(input: string) {
  return parseInput(input).filter(doShiftsOverlapEntirely).length;
}

export function part2(input: string) {
  return parseInput(input).filter(doShiftsOverlapABit).length;
}

type Shift = {
  start: number;
  end: number;
};

function parseInput(input: string): Shift[][] {
  return input
    .trim()
    .split("\n")
    .map((line) =>
      line.split(",").map((shift) => {
        const [start, end] = shift.split("-").map((hour) => parseInt(hour));
        return { start, end };
      })
    );
}

function doShiftsOverlapEntirely([s1, s2]: Shift[]): boolean {
  const firstOverlapsSecond = s1.start >= s2.start && s1.end <= s2.end;
  const secondOverlapsFirst = s2.start >= s1.start && s2.end <= s1.end;
  return firstOverlapsSecond || secondOverlapsFirst;
}

function doShiftsOverlapABit([s1, s2]: Shift[]): boolean {
  return !(s1.end < s2.start || s1.start > s2.end);
}
