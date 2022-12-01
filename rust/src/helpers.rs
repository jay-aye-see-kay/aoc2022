use std::fs;

/// helper fn to read the input file all in one
pub fn read_input(day: i32) -> String {
    let path = format!("../inputs/day{day:0>2}.txt");
    fs::read_to_string(&path).expect(&format!("Could not open {path}"))
}
