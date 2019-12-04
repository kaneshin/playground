fn f(x: f32) -> f32 {
    return x * x - x - 1.0;
}

fn midpoint_rule(a: f32, b: f32, n: i32) -> f32 {
    let h = (b - a) / n as f32;
    let mut r = 0.0;
    let mut x = a + h / 2.0;

    for _ in 0..n {
        r += f(x);
        x += h;
    }
    return r * h;
}

fn trapezoid_rule(a: f32, b: f32, n: i32) -> f32 {
    let h = (b - a) / n as f32;
    let mut r = 0.0;
    let mut x = a + h;
    let mut y1 = f(a);

    for _ in 0..n {
        let y2 = f(x);
        r += y1 + y2;
        x += h;
        y1 = y2;
    }
    return r * h / 2.0;
}

fn main() {
    println!("Midpoint Rule: {}", midpoint_rule(-2.0, 4.0, 1000));
    println!("Trapezoid Rule: {}", trapezoid_rule(-2.0, 4.0, 1000));
}
