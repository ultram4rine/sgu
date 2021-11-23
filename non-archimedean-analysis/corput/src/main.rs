use plotters::prelude::*;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let n = 25 as i64;
    let b = 2 as i64;

    // create graph picture and fill it by white color.
    let root = BitMapBackend::new("graph.png", (2000, 1000)).into_drawing_area();
    root.fill(&WHITE)?;
    let (left, right) = root.split_horizontally(1000);

    let make_chart = |area, from_x, to_x, from_y, to_y| {
        ChartBuilder::on(area)
            .caption("", ("sans-serif", 50).into_font())
            .margin(5)
            .x_label_area_size(30)
            .y_label_area_size(30)
            .build_cartesian_2d(from_x..to_x, from_y..to_y)
    };
    let mut chart_left = make_chart(&left, -0.1, 1.1, -0.1, 1.1)?;
    chart_left.configure_mesh().draw()?;
    let mut chart_right = make_chart(&right, -0.1_f64, n as f64 + 0.1, -0.1, 1.1)?;
    chart_right.configure_mesh().draw()?;

    for k in 0..=n {
        chart_left.draw_series(PointSeries::of_element(
            ((1_f64 - 1_f64 / b.pow(k as u32) as f64)
                ..(1_f64 - 1_f64 / b.pow((k + 1) as u32) as f64))
                .step(1_f64 / (10 * n) as f64)
                .values()
                .map(|x| (x, T(x, b, k as u32))),
            1,
            &RED,
            &|c, s, st| EmptyElement::at(c) + Circle::new((0, 0), s, st.filled()),
        ))?;
    }

    chart_left
        .draw_series(PointSeries::of_element(
            (0..0).map(|x| (x as f64, x as f64)),
            1,
            &RED,
            &|c, _s, _st| EmptyElement::at(c),
        ))?
        .label("Kakutani - von Neumann")
        .legend(|(x, y)| PathElement::new(vec![(x, y), (x + 20, y)], &RED));

    chart_right
        .draw_series(LineSeries::new(
            (0_i64..=n).map(|x| (x as f64, phi(x))),
            &BLUE,
        ))?
        .label("Monna")
        .legend(|(x, y)| PathElement::new(vec![(x, y), (x + 20, y)], &BLUE));

    chart_left
        .configure_series_labels()
        .background_style(&WHITE.mix(0.8))
        .border_style(&BLACK)
        .draw()?;
    chart_right
        .configure_series_labels()
        .background_style(&WHITE.mix(0.8))
        .border_style(&BLACK)
        .draw()?;

    Ok(())
}

#[allow(non_snake_case)]
/// Kakutani - von Neumann
fn T(x: f64, b: i64, k: u32) -> f64 {
    x - 1_f64 + (1_f64 / b.pow(k) as f64) + (1_f64 / b.pow(k + 1) as f64)
}

/// Monna
fn phi(x: i64) -> f64 {
    let binary = format!("{:b}", x).chars().rev().collect::<String>();

    let res = binary.chars().enumerate().fold(0_f64, |acc, (i, c)| {
        let xi = c.to_digit(2).unwrap();
        acc + xi as f64 * 2_f64.powf(-(i as f64 + 1_f64))
    });
    res
}
