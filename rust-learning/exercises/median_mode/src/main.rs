use std::collections::HashMap;

enum MiddleLength {
    Even,
    Odd,
}
impl MiddleLength {
    fn get_middle(numbers: &Vec<i32>) -> MiddleLength {
        if numbers.len() % 2 == 0 {
            MiddleLength::Even
        } else {
            MiddleLength::Odd
        }
    }
}

fn main() {
    let mut numbers = vec![37, 4, 8, 25, 1, 48, 7, 14, 7, 8, 7];

    numbers.sort_unstable();
    println!("The sorted is {:?}", numbers);

    let median = get_median(&numbers);

    println!("The median is: {median}");

    let mode = get_mode(&numbers);
    println!("The mode is: {mode}");
}

fn get_median(numbers: &Vec<i32>) -> i32 {
    let median_choice = MiddleLength::get_middle(numbers);
    match median_choice {
        MiddleLength::Odd => {
            let middle_number = numbers.len() / 2;
            *numbers.get(middle_number).unwrap()
        }
        MiddleLength::Even => {
            let middle_number = numbers.len() / 2;
            (numbers.get(middle_number).unwrap() + numbers.get(middle_number - 1).unwrap()) / 2
        }
    }
}

fn get_mode(numbers: &Vec<i32>) -> i32 {
    let mut mode_count = HashMap::new();

    for num in numbers {
        let count = mode_count.entry(num).or_insert(1);
        *count += 1;
    }

    let mut mode_key = -1;
    for (key, value) in &mode_count {
        if value > mode_count.get(&mode_key).unwrap_or(&1) {
            mode_key = **key;
        }
    }

    mode_key
}
