use rand::Rng;
use std::cmp::Ordering;
use std::io;

fn main() {
    println!("Guess the number!");

    let secret_number = rand::thread_rng().gen_range(1..=100);

    let mut input_message = "Please input your guess";

    loop {
        // println!("Please input your guess.");
        println!("{input_message}");

        let mut guess = String::new();

        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");

        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => {
                // println!("Invalid number. Try again.");
                input_message = "Invalid number. Try again.";
                continue;
            }
        };

        input_message = "Please input your guess";

        println!("You guessed: {guess}");

        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }
    // println!("You guessed: {}", 2 + 2);
}
