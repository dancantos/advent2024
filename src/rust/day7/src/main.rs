use std::time::Instant;

mod puzzle1;
mod input;

fn main() {
    let equations = input::read_input("input.txt");

    let now = Instant::now();
    let solution = puzzle1::sum_solveable(&equations);
    let elapsed = now.elapsed();
    println!("{}", solution);
    println!("elapsed: {:.2?}", elapsed);
    // println!("{}", solveable(10, vec![1, 9]));
}
