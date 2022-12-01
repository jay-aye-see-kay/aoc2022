use crate::helpers::read_input;
use std::fs::read_to_string;

fn get_cal_totals() -> Vec<i32> {
    let mut cal_totals: Vec<_> = read_input(1)
        .trim()
        .split("\n\n")
        .map(|group| {
            group
                .split('\n')
                .map(|line| line.parse::<i32>().unwrap())
                .sum()
        })
        .collect();
    cal_totals.sort_unstable();
    cal_totals.reverse();

    cal_totals
}

fn part1() -> i32 {
    get_cal_totals()[0]
}

fn part2() -> i32 {
    get_cal_totals()[..3].iter().sum()
}

mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1(), 71934)
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(), 211447)
    }
}
