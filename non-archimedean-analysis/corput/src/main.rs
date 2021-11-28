#![allow(non_snake_case)]

use plotters::prelude::*;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let n = 100 as i64;
    let b = 2 as i64;

    let name = format!("graphs_n_{}.png", n);

    // create graph picture and fill it by white color.
    let root = BitMapBackend::new(&name, (2000, 1000)).into_drawing_area();
    root.fill(&WHITE)?;
    let (left, right) = root.split_horizontally(1000);

    let make_chart = |area, (from_x, to_x), (from_y, to_y)| {
        ChartBuilder::on(area)
            .caption("", ("sans-serif", 50).into_font())
            .margin(5)
            .x_label_area_size(30)
            .y_label_area_size(30)
            .build_cartesian_2d(from_x..to_x, from_y..to_y)
    };
    let mut chart_left = make_chart(&left, (-0.1, 1.1), (-0.1, 1.1))?;
    chart_left.configure_mesh().draw()?;
    let mut chart_right = make_chart(&right, (-0.1_f64, (n - 1) as f64 + 0.1), (-0.1, 1.1))?;
    chart_right.configure_mesh().draw()?;

    let vdc_seq_R = vdc_R(b, n);
    chart_left
        .draw_series(PointSeries::of_element(
            (0..n).map(|i| (vdc_seq_R[i as usize], vdc_seq_R[i as usize + 1])),
            2,
            &RED,
            &|c, s, st| EmptyElement::at(c) + Circle::new((0, 0), s, st.filled()),
        ))?
        .label("Kakutani - von Neumann")
        .legend(|(x, y)| PathElement::new(vec![(x, y), (x + 20, y)], &RED));

    let vdc_seq_Z2 = vdc_Z2(n);
    chart_right
        .draw_series(LineSeries::new(
            (0..n).map(|i| {
                let x = vdc_seq_Z2[i as usize];
                (x as f64, phi(x))
            }),
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

/// Build the van der Corput sequence from Kakutani - von Neumann transformation.
fn vdc_R(b: i64, n: i64) -> Vec<f64> {
    let mut x = 0_f64;
    let mut seq = vec![];
    for _ in 0..=n {
        seq.push(x);

        let mut k = 0;
        while x < 1_f64 - (1_f64 / b.pow(k) as f64) || x >= 1_f64 - (1_f64 / b.pow(k + 1) as f64) {
            k += 1;
        }

        x = T(x, b, k);
    }
    seq
}

/// Build the van der Corput sequence from Kakutani - von Neumann transformation.
fn vdc_Z2(n: i64) -> Vec<i64> {
    let mut x = 0_i64;
    let mut seq = vec![];
    for _ in 0..=n {
        seq.push(x);
        x = tau(x);
    }
    seq
}

/// Kakutani - von Neumann
fn T(x: f64, b: i64, k: u32) -> f64 {
    x - 1_f64 + (1_f64 / b.pow(k) as f64) + (1_f64 / b.pow(k + 1) as f64)
}

/// Z2 analogue of Kakutani - von Neumann mapping
fn tau(x: i64) -> i64 {
    x + 1
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
