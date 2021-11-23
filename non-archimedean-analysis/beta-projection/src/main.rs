use plotters::prelude::*;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    // create graph picture and fill it by white color.
    let root = BitMapBackend::new("graph.png", (1000, 1000)).into_drawing_area();
    root.fill(&WHITE)?;
    // create area for graph.
    let mut chart = ChartBuilder::on(&root)
        .caption("", ("sans-serif", 50).into_font())
        .margin(5)
        .x_label_area_size(30)
        .y_label_area_size(30)
        .build_cartesian_2d(-0.1..1.1, -0.1..1.1)?;

    chart.configure_mesh().draw()?;

    let n = 20 as f64;
    let beta = 2_f64.powf(1_f64 / n);
    let k = 15;
    // LineSeries connects separate points, so graph becomes incorrect.
    // Solution is to use PointSeries with points as circles with size 1.
    chart
        .draw_series(PointSeries::of_element(
            (0..2_i64.pow(k)).map(|x| e(k, x, beta)),
            1,
            &RED,
            &|c, s, st| EmptyElement::at(c) + Circle::new((0, 0), s, st.filled()),
        ))?
        .label(format!("f(x) with k = {}", k))
        .legend(|(x, y)| PathElement::new(vec![(x, y), (x + 20, y)], &RED));

    chart
        .configure_series_labels()
        .background_style(&WHITE.mix(0.8))
        .border_style(&BLACK)
        .draw()?;

    Ok(())
}

fn f(x: i64) -> i64 {
    // test function
    // x + 1
    (x ^ 1)
        ^ (2 * (x
            & (1 + 2 * x)
            & (3 + 4 * x)
            & (7 + 8 * x)
            & (15 + 16 * x)
            & (31 + 32 * x)
            & (63 + 64 * x)))
        ^ (4 * (x * x + 8))
}

fn e(k: u32, x: i64, beta: f64) -> (f64, f64) {
    let beta_projection = |x: i64| {
        let binary = format!("{:b}", x).chars().rev().collect::<String>();
        let numerator = binary.chars().enumerate().fold(0_f64, |acc, (i, c)| {
            let xi = c.to_digit(2).unwrap();
            acc + xi as f64 * beta.powf(i as f64)
        });
        (numerator / beta.powf(k as f64)) % 1_f64
    };
    (beta_projection(x), beta_projection(f(x)))
}
