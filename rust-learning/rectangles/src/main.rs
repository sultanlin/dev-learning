#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }

    fn can_hold(&self, other_rect: &Rectangle) -> bool {
        self.width >= other_rect.width && self.height >= other_rect.height
    }

    fn square(size: u32) -> Self {
        Self {
            width: size,
            height: size,
        }
    }
}

fn main() {
    let scale = 2;
    let my_rectangle = Rectangle {
        width: dbg!(30 * scale),
        height: 50,
    };

    // dbg!(&my_rectangle);

    // println!("My rectangle is {:#?}", my_rectangle);

    // println!("My rectangle's width is: {}", my_rectangle.width);
    // // println!("My rectangle's height is: {}", my_rectangle.height);
    //
    println!(
        "The area of the rectangle is {} square pixels.",
        // area(width1, height1)
        // area(&my_rectangle)
        my_rectangle.area()
    );

    let my_square = Rectangle::square(30);

    println!(
        "The area of the rectangle is {} square pixels.",
        // area(width1, height1)
        // area(&my_rectangle)
        my_square.area()
    );
}
