use std::{borrow::BorrowMut, cell::RefCell, ops::Deref, rc::Rc};

use crate::List::{Cons, Nil};

// enum List {
//     Cons(i32, Box<List>),
//     Nil,
// }

// enum List {
//     Cons(i32, Rc<List>),
//     Nil,
// }

#[derive(Debug)]
enum List {
    Cons(Rc<RefCell<i32>>, Rc<List>),
    Nil,
}

struct MyBox<T>(T);

impl<T> MyBox<T> {
    fn new(x: T) -> MyBox<T> {
        MyBox(x)
    }
}

impl<T> Deref for MyBox<T> {
    type Target = T;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

// impl<T> Drop for MyBox<T> {
//     // fn drop(&self)
// }

struct CustomSmartPointer {
    data: String,
}

impl Drop for CustomSmartPointer {
    fn drop(&mut self) {
        println!("Dropping CustomSmartPointer with data '{}'!", self.data);
    }
}

fn main() {
    // let list = Cons(1, Box::new(Cons(2, Box::new(Cons(3, Box::new(Nil))))));

    // let x = 5;
    // // let y = &x;
    // let y = MyBox::new(x);
    //
    // assert_eq!(5, x);
    // assert_eq!(5, *y);
    // // println!("{}", assert_eq!(5, *y));
    //
    // let m = MyBox(String::from("Rust"));
    // hello(&m);

    // let c = CustomSmartPointer {
    //     data: String::from("my stuff"),
    // };
    // let d = CustomSmartPointer {
    //     data: String::from("other stuff"),
    // };
    // drop(c);
    // println!("CustomSmartPointer created.");

    // let a = Rc::new(Cons(5, Rc::new(Cons(10, Rc::new(Nil)))));
    // println!("count after creating a = {}", Rc::strong_count(&a));
    // let b = Cons(3, Rc::clone(&a));
    // println!("count after creating b = {}", Rc::strong_count(&a));
    // {
    //     let c = Cons(4, Rc::clone(&a));
    //     println!("count after creating c = {}", Rc::strong_count(&a));
    // }
    // println!("count after c goes out of scope = {}", Rc::strong_count(&a));

    // let a = rc::new(cons(5, rc::new(cons(10, rc::new(nil)))));
    // println!("count after creating a = {}", rc::strong_count(&a));
    // let b = cons(3, rc::clone(&a));
    // println!("count after creating b = {}", rc::strong_count(&a));
    // {
    //     let c = cons(4, rc::clone(&a));
    //     println!("count after creating c = {}", rc::strong_count(&a));
    // }
    // println!("count after c goes out of scope = {}", rc::strong_count(&a));

    let value = Rc::new(RefCell::new(5));

    let a = Rc::new(Cons(Rc::clone(&value), Rc::new(Nil)));

    let b = Cons(Rc::new(RefCell::new(3)), Rc::clone(&a));
    let c = Cons(Rc::new(RefCell::new(4)), Rc::clone(&a));

    *value.borrow_mut() += 10;

    println!("a after = {:?}", a);
    println!("b after = {:?}", b);
    println!("c after = {:?}", c);
}

fn hello(name: &str) {
    println!("Hello, {name}!");
}
