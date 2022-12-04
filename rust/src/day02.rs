use crate::helpers::read_input;
use std::fs::read_to_string;

fn part1() -> i32 {
    read_input(2)
        .lines()
        .map(|line| match line {
            "A X" => 4,
            "A Y" => 8,
            "A Z" => 3,
            "B X" => 1,
            "B Y" => 5,
            "B Z" => 9,
            "C X" => 7,
            "C Y" => 2,
            "C Z" => 6,
            _ => panic!("unknown line"),
        })
        .sum()
}

#[derive(Debug)]
enum Choice {
    Rock,
    Paper,
    Scissors,
}
impl Choice {
    fn points(self) -> i32 {
        match self {
            Self::Rock => 1,
            Self::Paper => 2,
            Self::Scissors => 3,
        }
    }
}

fn part2() -> i32 {
    read_input(2)
        .lines()
        .map(|line| {
            let my_choice = match line {
                "A X" | "B Z" | "C Y" => Choice::Scissors,
                "A Y" | "B X" | "C Z" => Choice::Rock,
                "A Z" | "B Y" | "C X" => Choice::Paper,
                _ => panic!("unknown line"),
            };
            let result_points = match line.chars().nth(2).unwrap() {
                'X' => 0,
                'Y' => 3,
                'Z' => 6,
                _ => panic!("unknown line"),
            };
            my_choice.points() + result_points
        })
        .sum()
}

mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1(), 11906)
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(), 11186)
    }
}
