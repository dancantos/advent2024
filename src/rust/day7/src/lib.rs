
#![feature(test)]
extern crate test;

pub mod puzzle1;
pub mod input;

#[cfg(test)]
mod tests {
    use super::*;
    use test::Bencher;

    // 50,944.36 ns/iter (+/- 16,467.42)
    #[bench]
    fn bench_sum_solveable(b: &mut Bencher) {
        let equations = input::read_input("input.txt");
        b.iter(|| puzzle1::sum_solveable(&equations));
    }
}
