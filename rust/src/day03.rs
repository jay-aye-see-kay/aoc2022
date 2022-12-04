use crate::helpers::read_input;
use std::{collections::HashSet, fs::read_to_string};

fn part1() -> i32 {
    read_input(3)
        .lines()
        .map(|line| {
            let (left, right) = line.split_at(line.len() / 2);
            let left = left.chars().collect::<HashSet<char>>();
            let right = right.chars().collect::<HashSet<char>>();
            score(left.intersection(&right).next().unwrap())
        })
        .sum()
}

fn part2() -> i32 {
    read_input(3)
        .lines()
        .collect::<Vec<_>>()
        .chunks(3)
        .map(|lines| {
            let sets = lines
                .iter()
                .map(|line| line.chars().collect::<HashSet<char>>())
                .reduce(|accum, set| accum.intersection(&set).cloned().collect())
                .unwrap();
            score(sets.iter().next().unwrap())
        })
        .sum()
}

fn score(c: &char) -> i32 {
    let char_code = *c as i32;
    match char_code {
        65..=90 => char_code - 38,
        97..=122 => char_code - 96,
        _ => panic!("unknown char"),
    }
}

mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1(), 7428)
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(), 2650)
    }
}
