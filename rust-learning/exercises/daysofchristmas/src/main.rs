fn main() {
    let num_to_english = [
        "first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eightth", "ninth",
        "tenth", "eleventh", "twelveth",
    ];
    let received = [
        "a partridge in a pear tree",
        "Two turtle doves",
        "Three French hens",
        "Four calling birds",
        "Five gold rings",
        "Six geese a-laying",
        "Seven swans a-swimming",
        "Eight maids a-milking",
        "Nine ladies dancing",
        "Ten lords a-leaping",
        "Eleven pipers piping",
        "Twelve drummers drumming",
    ];

    for day in 0..num_to_english.len() {
        println!(
            "In the {} day of Christmas my true love sent to me",
            num_to_english[day]
        );
        if day == 0 {
            println!("{}.\n", received[day]);
            continue;
        }
        for gift in (0..day + 1).rev() {
            if gift == 0 {
                println!("And {}.\n", received[gift]);
                continue;
            }
            // println!("{gift}, ");
            println!("{},", received[gift])
        }
    }
}
