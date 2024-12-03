fn main() {
    let a: [i64; 5] = [1, 2, 3, 4, 5];
    let b: [i64; 5] = [2, 2, 3, 4, 5];
    let s = sumdiffs(&a, &b);
    println!("{s}");
}

fn sumdiffs(a: &[i64], b: &[i64]) -> i64 {
    let mut sum: i64 = 0;
    for i in 0..a.len() {
        sum += abs(a.get(i).unwrap() - b.get(i).unwrap())
    }
    return sum
}

fn abs(n: i64) -> i64 {
    if n < 0 {
        return -n;
    }
    return n;
}
