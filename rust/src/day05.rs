use crate::helpers::read_input;
use std::{collections::HashSet, fs::read_to_string};

fn part1() -> String {
    let (mut stacks, moves) = parse_input(read_input(5));
    for mv in moves {
        for _ in 0..mv.count {
            let from = stacks[mv.from].items.pop().unwrap();
            stacks[mv.to].items.push(from);
        }
    }
    stacks
        .iter()
        .map(|stack| stack.items.last().unwrap())
        .collect()
}

fn part2() -> String {
    let (mut stacks, moves) = parse_input(read_input(5));
    for mv in moves {
        let mut tmp: Vec<_> = (0..mv.count)
            .map(|_| stacks[mv.from].items.pop().unwrap())
            .collect();

        (0..mv.count).for_each(|_| {
            stacks[mv.to].items.push(tmp.pop().unwrap());
        });
    }
    stacks
        .iter()
        .map(|stack| stack.items.last().unwrap())
        .collect()
}

#[derive(Debug)]
struct Move {
    count: usize,
    from: usize,
    to: usize,
}

#[derive(Debug)]
struct Stack {
    items: Vec<char>,
}

fn parse_input(input: String) -> (Vec<Stack>, Vec<Move>) {
    let stacks: Vec<Stack> = [
        "SPHVFG", "MZDVBFJG", "NJLMG", "PWDVZGN", "BCRV", "ZLWPMSRV", "PHT", "VZHCNSRQ", "JQVPGLF",
    ]
    .iter()
    .map(|str| Stack {
        items: str.chars().rev().collect(),
    })
    .collect();

    let (stacks_str, moves_str) = input.split_once("\n\n").unwrap();

    let moves: Vec<Move> = moves_str
        .split("\n")
        .map(|line| {
            let values: Vec<usize> = line
                .split(" ")
                .filter_map(|word| word.parse().ok())
                .collect();
            return Move {
                count: values[0],
                from: values[1] - 1,
                to: values[2] - 1,
            };
        })
        .collect();

    (stacks, moves)
}

mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1(), "FCVRLMVQP")
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(), "RWLWGJGFD")
    }
}
