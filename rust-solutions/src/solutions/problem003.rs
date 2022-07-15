use crate::printing;

pub fn solution() {
    printing::pretty_print_solution(compute, 3);
}


fn compute() -> i64 {
    let mut prime_factors = factors(&600851475143)
        .into_iter()
        .filter(|&x| is_prime(&x))
        .collect::<Vec<i64>>();

    prime_factors.sort();
    prime_factors[prime_factors.len() - 1]
}

fn factors(n: &i64) -> Vec<i64>{
    let mut res_vec: Vec<i64> = Vec::new();
    for i in 1..=((*n as f64).sqrt() as i64) {
        if (*n)%i == 0 {
            res_vec.push(i);
            res_vec.push((*n)/i as i64);
        }
    }

    res_vec
}

fn is_prime(n: &i64) -> bool {
    factors(n).len() == 2
}
