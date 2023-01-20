use crate::printing;

pub fn solution() {
    printing::pretty_print_solution(compute, 10);
}

fn compute() -> usize {
    let primes = prime_sieve(2_000_000);
    let mut answer = 0;
    for i in primes { answer += i; }

    answer
}

fn prime_sieve(n: usize) -> Vec<usize> {
    let mut primes: Vec<bool> = (0..n + 1).map(|num| num == 2 || num & 1 != 0).collect();
    let mut num = 3_usize;

    while num * num <= n {
        let mut j = num * num;
        while j <= n {
            primes[j] = false;
            j += num;
        }
        num += 2;
    }

    primes
        .into_iter()
        .enumerate()
        .skip(2)
        .filter_map(|(i, p)| if p { Some(i) } else { None })
        .collect::<Vec<usize>>()
}
