use std::io;
fn main() {
    println!("How do you want the conversion?");
    println!("A) Fahrenheit to Celsius.");
    println!("B) Celsius to Fahrenheit.");

    loop {
        let mut conversion_decision = String::new();
        io::stdin()
            .read_line(&mut conversion_decision)
            .expect("Failed to read line");

        let conversion_decision = conversion_decision.trim();
        if conversion_decision == "A" {
            println!("Fahrenheit to Celsius");
            let degrees = get_number();
            println!("The answer is: {:.2}.", fahrenheit_to_celsius(degrees));
            break;
        } else if conversion_decision == "B" {
            println!("Celsius to Fahrenheit");
            let degrees = get_number();
            println!("The answer is: {:.2}.", celsius_to_fahrenheit(degrees));
            break;
        } else {
            println!("Please pick A or B");
        }
    }

    println!("Hello, world!");
}

fn celsius_to_fahrenheit(x: f64) -> f64 {
    return (x * 9.0 / 5.0) + 32.0;
}
fn fahrenheit_to_celsius(x: f64) -> f64 {
    return (x - 32.0) * 5.0 / 9.0;
}

fn get_number() -> f64 {
    loop {
        println!("What is the current degree?");
        let mut degrees = String::new();
        io::stdin()
            .read_line(&mut degrees)
            .expect("Failed to read line");

        let degrees: f64 = match degrees.trim().parse() {
            Ok(num) => num,
            Err(_) => {
                print!("Enter valid number. ");
                continue;
            }
        };

        return degrees;
    }
}
