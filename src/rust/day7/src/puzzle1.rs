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
    if target % last == 0 && _solveable(target/last, terms, head-1) {
        return true
    }
    return _solveable(target-last, terms, head-1)
}
