struct Interval(f32, f32);

fn f(x: f32) -> f32 {
    return x * x - x - 1.0;
}

fn midpoint_rule(i: Interval, n: i32) -> f32 {
    let h = (i.1 - i.0) / n as f32;
    let mut r = 0.0;
    let mut x = i.0 + h / 2.0;

    for _ in 0..n {
        r += f(x);
        x += h;
    }
    return r * h;
}

fn trapezoid_rule(i: Interval, n: i32) -> f32 {
    let h = (i.1 - i.0) / n as f32;
    let mut r = 0.0;
    let mut x = i.0 + h;
    let mut y1 = f(i.0);

    for _ in 0..n {
        let y2 = f(x);
        r += y1 + y2;
        x += h;
        y1 = y2;
    }
    return r * h / 2.0;
}

fn main() {
    println!("Midpoint Rule: {}", midpoint_rule(Interval(-2.0, 4.0), 1000));
    println!("Trapezoid Rule: {}", trapezoid_rule(Interval(-2.0, 4.0), 1000));
}
