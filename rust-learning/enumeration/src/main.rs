#[derive(Debug)]
enum UsState {
    Alabama,
    Alaska,
    Carolina,
}

enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter(UsState),
}

fn main() {
    println!("{}", value_in_cents(Coin::Quarter(UsState::Carolina)));

    let some_num = Some(5);
    let y: Option<i32> = None;

    let x = plus_one(some_num);

    let z = Some(100);

    if let Some(i) = z {
        println!("{i}")
    } else {
        println!("Help!");
    };

    // println!("{summm}")
}

fn plus_one(x: Option<i32>) -> Option<i32> {
    match x {
        None => None,
        Some(i) => Some(i + 1),
    }
}

fn value_in_cents(coin: Coin) -> u8 {
    match coin {
        Coin::Penny => 1,
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quarter(state) => {
            println!("State quarter from {:?}!", state);
            25
        }
    }
}
// enum Message {
//     Quit,
//     Move { x: i32, y: i32 },
//     Write(String),
//     ChangeColor(i32, i32, i32),
// }
//
// impl Message {
//     fn call(&self) {}
// }
//
// fn main() {
//     let m = Message::Write(String::from("hello"));
//     m.call()
// }
// // enum IpAddrKind {
// //     V4,
// //     V6,
// // }
//
// enum IpAddr {
//     V4(u8, u8, u8, u8),
//     // V4(String),
//     V6(String),
// }
//
// fn main() {
//     println!("Hello, world!");
//     //
//     // let four = IpAddrKind::V4;
//     // let six = IpAddrKind::V6;
//
//     // let home = IpAddr::V4(String::from("127.0.0.1"));
//     let home = IpAddr::V4(127, 0, 0, 1);
//     let loopback = IpAddr::V6(String::from("::1"));
// }
