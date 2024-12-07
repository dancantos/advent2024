use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

pub fn read_input(file: &str) -> Vec<(i64, Vec<i64>)> {
    let mut result = Vec::<(i64, Vec<i64>)>::new();
    if let Ok(lines) = read_lines(file) {
        for line in lines.flatten() {
            // split line into equation
            let mut split = line.split_whitespace();
            let mut target = split.next().unwrap().to_string();

            // first entry: rm trailing ':' and store as 'target'
            target.pop();
            let target_int = target.parse::<i64>().unwrap();

            // remainder: store as terms
            let mut terms = Vec::<i64>::new();
            for term in split {
                terms.push(term.parse::<i64>().unwrap());
            }
            result.push((target_int, terms));
        }
    }
    return result
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
