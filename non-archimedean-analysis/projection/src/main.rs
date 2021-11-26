use plotters::prelude::*;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let k = 20;

    let name = format!("k_{}.png", k);
    // create graph picture and fill it by white color.
    let root = BitMapBackend::new(&name, (1000, 1000)).into_drawing_area();
    root.fill(&WHITE)?;
    // create area for graph.
    let mut chart = ChartBuilder::on(&root)
        .caption("", ("sans-serif", 50).into_font())
        .margin(5)
        .x_label_area_size(30)
        .y_label_area_size(30)
        .build_cartesian_2d(-0.1..1.1, -0.1..1.1)?;

    chart.configure_mesh().draw()?;

    // LineSeries connects separate points, so graph becomes incorrect.
    // Solution is to use PointSeries with points as circles with size 1.
    chart
        .draw_series(PointSeries::of_element(
            (0..2_i64.pow(k)).map(|x| e(k, x)),
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
    //x + ((x * x) | 74071)
    (x ^ 1)
        ^ (2 * (x
            & (1 + 2 * x)
            & (3 + 4 * x)
            & (7 + 8 * x)
            & (15 + 16 * x)
            & (31 + 32 * x)
            & (63 + 64 * x)))
        ^ (4 * (x * x + 27))
}

fn e(k: u32, x: i64) -> (f64, f64) {
    let projection = |x: i64| (x % 2_i64.pow(k)) as f64 / 2_f64.powf(k as f64);
    (projection(x), projection(f(x)))
}
