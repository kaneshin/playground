fn print_fizzbuzz(n: i32) {
    for i in 1..(n + 1) {
        if i % 15 == 0 {
            println!("fizzbuzz");
        } else if i % 3 == 0 {
            println!("fizz");
        } else if i % 5 == 0 {
            println!("buzz");
        } else {
            println!("{}", i);
        }
    }
}

trait FizzBuzz {
    fn fizzbuzz(&self);
}

impl FizzBuzz for i32 {
    fn fizzbuzz(&self) {
        print_fizzbuzz(*self);
    }
}

impl FizzBuzz for i64 {
    fn fizzbuzz(&self) {
        print_fizzbuzz(*self as i32);
    }
}

fn main() {
    (20 as i32).fizzbuzz();
    (20 as i64).fizzbuzz();
}
