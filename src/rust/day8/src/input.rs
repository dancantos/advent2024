use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::collections::HashMap;

pub fn read_input(file: &str) -> (HashMap<char, Vec<(u64, u64)>>, (u64, u64)) {
    let mut result: HashMap<char, Vec<(u64, u64)>> = HashMap::new();
    let mut width: u64 = 0;
    let mut y: u64 = 0;
    if let Ok(lines) = read_lines(file) {
        for line in lines.flatten() {
            width = line.len().try_into().unwrap();
            let mut x: u64 = 0;
            for c in line.chars() {
                if c != '.' {
                    if result.contains_key(&c) {
                        result.get_mut(&c).unwrap().push((x, y));
                    } else {
                        let points = vec![(x, y)];
                        result.insert(c, points);
                    }
                }
                x += 1;
            }
            y += 1
        }
    }
    return (result, (width, y))
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
