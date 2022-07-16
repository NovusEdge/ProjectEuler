use crate::printing;
use itertools::Itertools;

pub fn solution() {
    printing::pretty_print_solution(compute, 4);
}


fn compute() -> i32 {   
    (100..1000)
        .cartesian_product((100..1000).rev())
        .map(|x| x.0 * x.1)
        .filter(|&x| is_pallindrome(x))
        .max()
        .unwrap()
}

fn is_pallindrome(n: i32) -> bool {
    let num_str: String = format!("{:?}", n);
    let reversed = num_str
        .chars()
        .rev()
        .collect::<String>();

    num_str == reversed
}
