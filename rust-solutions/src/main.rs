pub mod printing;
pub mod solutions;

use std::io::{stdin, stdout, Write};

fn main() {
    println!("Project Euler Solutions with rust!");
    print!("Enter Problem number: ");
    let _ = stdout().flush();

    let mut input = String::new();
    stdin()
        .read_line(&mut input)
        .ok()
        .expect("Failed to read input");
    let problem_index: i32 = input.trim().parse().expect("Invalid Problem Number");

    match problem_index {
        1  => solutions::problem001::solution(),
        2  => solutions::problem002::solution(),
        3  => solutions::problem003::solution(),
        4  => solutions::problem004::solution(),
        5  => solutions::problem005::solution(),
        6  => solutions::problem006::solution(),
        7  => solutions::problem007::solution(),
        8  => solutions::problem008::solution(),
        9  => solutions::problem009::solution(),
        10 => solutions::problem010::solution(),
        _  => println!("Problem {} not yet solved :(", problem_index),
    }
}
