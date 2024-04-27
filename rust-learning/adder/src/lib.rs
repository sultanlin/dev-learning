//! # My Crate
//!
//! `my_crate` is a collection of utilities to make performing certain
//! calculations more convenient.

#[cfg(test)]
mod tests {
    use super::*;

    // #[test]
    // fn exploration() {
    //     assert_eq!(2 + 2, 4);
    // }
    //
    // #[test]
    // fn another() {
    //     panic!("Make this test fail!");
    // }
    //
    // #[test]
    // fn larger_can_hold_smaller() {
    //     let larger = Rectangle {
    //         width: 8,
    //         height: 7,
    //     };
    //     let smaller = Rectangle {
    //         width: 5,
    //         height: 1,
    //     };
    //     assert!(larger.can_hold(&smaller));
    // }
    //
    // #[test]
    // fn smaller_cannot_hold_larger() {
    //     let larger = Rectangle {
    //         width: 8,
    //         height: 7,
    //     };
    //     let smaller = Rectangle {
    //         width: 5,
    //         height: 1,
    //     };
    //     assert!(!smaller.can_hold(&larger));
    // }
    //
    // #[test]
    // fn it_adds_two() {
    //     assert_eq!(4, add_two(2));
    // }
    //
    // #[test]
    // fn greeting_contains_name() {
    //     let result = greeting("Carol");
    //     assert!(
    //         result.contains("Carol"),
    //         "Greeting did not contain name, value was '{}'",
    //         result
    //     );
    // }
    //
    // #[test]
    // fn it_works() -> Result<(), String> {
    //     if 2 + 2 == 4 {
    //         Ok(())
    //     } else {
    //         Err(String::from("two plus two does not equal four"))
    //     }
    // }
    //
    // #[test]
    // #[ignore]
    // fn add_two_and_two() {
    //     assert_eq!(4, add_two(2));
    // }
    //
    // #[test]
    // fn add_three_and_two() {
    //     assert_eq!(5, add_two(3));
    // }
    //
    #[test]
    fn one_hundred() {
        assert_eq!(102, add_two(100));
    }
}

#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width < other.width && self.height > other.height
    }
}

/// Adds two to the number given.
///
/// # Examples
///
/// ```
/// let arg = 5;
/// let answer = adder::add_two(arg);
///
/// assert_eq!(7, answer);
/// ```
pub fn add_two(a: i32) -> i32 {
    a + 2
}

pub fn add_one(a: i32) -> i32 {
    a + 1
}

pub fn greeting(name: &str) -> String {
    format!("Hello {}!", name)
}
