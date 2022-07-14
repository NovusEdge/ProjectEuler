use crate::printing;

pub fn solution() {
    printing::pretty_print_solution(compute, 2);
}


fn compute() -> i64 {
    let (mut sum, mut i) = (0_i64, 1_i32);

    loop {
        let term = fibbonacci_term(i);
        if term > 4_000_000 { 
            break 
        }

        if term%2 == 0 {
            sum += term;
        } 

        i += 1;
    }

    sum
}

fn fibbonacci_term(n: i32) -> i64 {
    match n {
        0 | 1 => 1,
        _ => fibbonacci_term(n-2) + fibbonacci_term(n-1),
    }
}
