use std::io;
use std::io::BufRead;

fn main() {
    let stdin = io::stdin();
    let reader = stdin.lock();

    let mut calibration_sum: u32 = 0;
    for line in reader.lines() {
        match line {
            Ok(input) => {
                calibration_sum += get_calibration_value(&input);
            }
            Err(_) => break
        }
    }
    
    println!("{}", calibration_sum)
}

fn get_calibration_value(s: &String) -> u32 {
    let mut first: u32 = 0;
    let mut last: u32 = 0;

    for c in s.chars() {
        if let Some(digit) = c.to_digit(10) {
            first = digit;
            break;
        }
    }
    for c in s.chars().rev() {
        if let Some(digit) = c.to_digit(10) {
            last = digit;
            break;
        }
    }
    10 * first + last
}