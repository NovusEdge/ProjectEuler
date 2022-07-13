use crate::printing;

pub fn solution() {
    printing::pretty_print_solution(compute, 1);
}

pub fn compute() -> i32 {
    let mut sum: i32 = 0;
    for i in 0..1000 {
        if i%3 == 0 || i%5 == 0 {
            sum += i;
        }
    }

    sum
}
