use crate::printing;

pub fn solution() {
    printing::pretty_print_solution(compute, 9);
}

fn compute() -> usize {
    for a in 1..500 {
        for b in 1..500 {
            let c = 1000 - (a + b);
            if a * a + b * b == c * c {
                return a * b * c;
            }
        }
    }

    0
}
