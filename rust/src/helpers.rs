use std::fs;

/// helper fn to read the input file all in one
pub fn read_input(day: i32) -> String {
    let path = format!("../inputs/day{day:0>2}.txt");
    fs::read_to_string(&path)
        .unwrap_or_else(|_| panic!("Could not open {path}"))
        .trim()
        .to_owned()
}
