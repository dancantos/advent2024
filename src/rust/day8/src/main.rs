use grid::Bitmask;

pub mod input;

fn main() {
    let (grid, size) = input::read_input("input.txt");
    for (k, v) in grid.iter() {
        println!("{}: {:?}", k, v);
    }
    let mut bm = Bitmask::new(size.0, size.1);
    bm.set(1, 1);
    println!("{}", bm.is_set(1, 1));
}


