use crate::printing;

pub fn solution() {
    printing::pretty_print_solution(compute, 6);
}

fn compute() -> i64 {
    (sum_of_squares(100) - (sum_of_numbers(100) * sum_of_numbers(100))).abs()
}

fn sum_of_squares(n: i64) -> i64 {
    n * (n + 1) * (2 * n + 1) / 6
}

fn sum_of_numbers(n: i64) -> i64 {
    n * (n + 1) / 2
}
