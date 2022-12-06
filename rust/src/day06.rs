use crate::helpers::read_input;
use std::collections::HashSet;

fn find_unique_substr(substr_len: usize) -> Option<usize> {
    let chars: Vec<_> = read_input(6).chars().collect();
    for (start_idx, window) in chars.windows(substr_len).enumerate() {
        let set: HashSet<_> = window.iter().collect();
        if set.len() == substr_len {
            return Some(start_idx + substr_len);
        }
    }
    return None;
}

fn part1() -> usize {
    find_unique_substr(4).expect("substring not found")
}

fn part2() -> usize {
    find_unique_substr(14).expect("substring not found")
}

mod test {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1(), 1816)
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(), 2625)
    }
}
