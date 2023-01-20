use crate::printing;
use num::integer::gcd;

pub fn solution() {
    printing::pretty_print_solution(compute, 5);
}

fn compute() -> i64 {
    (1..21).reduce(|a, b| (a * b) / gcd(a, b)).unwrap()
}
