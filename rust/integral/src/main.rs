trait Integral {
    fn method(&self) -> String;
    fn calc(&self) -> f32;
    fn print(&self) {
        println!("{}: {}", self.method(), self.calc());
    }
}

struct Interval(f32, f32);

struct MidpointRule {
    function: fn(f32) -> f32,
    interval: Interval,
    step: i32,
}

impl Integral for MidpointRule {
    fn method(&self) -> String {
        return String::from("Midpoint Rule");
    }

    fn calc(&self) -> f32 {
        let a = self.interval.0;
        let b = self.interval.1;
        let n = self.step;
        let h = (b - a) / n as f32;
        let mut r = 0.0;
        let mut x = a + h / 2.0;

        for _ in 0..n {
            r += (self.function)(x);
            x += h;
        }
        return r * h;
    }
}

struct TrapezoidRule {
    function: fn(f32) -> f32,
    interval: Interval,
    step: i32,
}

impl Integral for TrapezoidRule {
    fn method(&self) -> String {
        return String::from("Trapezoid Rule");
    }

    fn calc(&self) -> f32 {
        let a = self.interval.0;
        let b = self.interval.1;
        let n = self.step;
        let h = (b - a) / n as f32;
        let mut r = 0.0;
        let mut x = a + h;
        let mut y1 = (self.function)(a);

        for _ in 0..n {
            let y2 = (self.function)(x);
            r += y1 + y2;
            x += h;
            y1 = y2;
        }
        return r * h / 2.0;
    }
}

fn main() {
    fn f(x: f32) -> f32 {
        return x * x - x - 1.0;
    }

    MidpointRule {
        function: f,
        interval: Interval(-2.0f32, 4.0f32),
        step: 1000,
    }.print();

    TrapezoidRule {
        function: f,
        interval: Interval(-2.0f32, 4.0f32),
        step: 1000,
    }.print();
}
