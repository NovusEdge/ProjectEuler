pub fn pretty_print_solution<T: std::fmt::Display>(f: fn() -> T, problem_index: i32) {
    let now = std::time::Instant::now();
    let answer: T = f();
    let elapsed = now.elapsed();

    println!("\nAnswer to Problem {}: {}", problem_index, answer);
    println!("Time Taken: {:?}\n", elapsed);
}
