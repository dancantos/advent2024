pub fn sum_solveable(equations: &[(i64, Vec<i64>)]) -> i64 {
    let mut count = 0;
    for eq in equations {
        if solveable(eq.0, &eq.1) {
            count += eq.0;
        }
    }
    return count;
}

pub fn solveable(target: i64, terms: &[i64]) -> bool {
    return _solveable(target, terms, terms.len()-1)
}

fn _solveable(target: i64, terms: &[i64], head: usize) -> bool {
    let last = terms[head];
    if head == 0 {
        return target == last
    }
    // try dividing
    if target % last == 0 && _solveable(target/last, terms, head-1) {
        return true
    }
    // try unconcat
    let _unconcat = unconcat(target, last);
    if _unconcat < target && _solveable(_unconcat, terms, head-1) {
        return true
    }
    return _solveable(target-last, terms, head-1)
}

fn unconcat(a: i64, b: i64) -> i64 {
    if a == b {
        return 0
    }
    let order = (0.1 + (b as f64)).log10().ceil() as i64;
    let exp = 10i64.pow(order as u32);
    if a % exp == b {
        return (a - b) / exp
    }
    return a
}
