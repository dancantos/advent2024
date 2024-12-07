
#![feature(test)]
extern crate test;

pub mod puzzle1;
pub mod puzzle2;
pub mod input;

#[cfg(test)]
mod tests {
    use super::*;
    use test::Bencher;

    // 50,944.36 ns/iter (+/- 16,467.42)
    #[bench]
    fn bench_puzzle1(b: &mut Bencher) {
        let equations = input::read_input("input.txt");
        b.iter(|| puzzle1::sum_solveable(&equations));
    }

    // 207,162.24 ns/iter (+/- 35,680.19)
    #[bench]
    fn bench_puzzle2(b: &mut Bencher) {
        let equations = input::read_input("input.txt");
        b.iter(|| puzzle2::sum_solveable(&equations));
    }
}
