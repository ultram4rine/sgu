#![allow(non_snake_case)]

use plotters::prelude::*;
use rand::prelude::*;
use std::fs;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let m: usize = 8;

    let H = f;

    // Graph stuff.
    let name = &format!("graph_m_{}.png", m);
    let root = BitMapBackend::new(&name, (1000, 1000)).into_drawing_area();
    root.fill(&WHITE)?;

    let axis = -30..2_i64.pow(m as u32) + 30;
    let mut chart = ChartBuilder::on(&root).margin(30).build_cartesian_3d(
        axis.clone(),
        axis.clone(),
        2_i64.pow(m as u32) + 30..-30,
    )?;

    chart.configure_axes().draw()?;

    // Not easy to understand but it actually a nested loop in iterator form.
    // Alternative to `for x in xs { for z in zs { f(x, z) } }`.
    let points = (0..2_i64.pow(m as u32))
        .map(|x| (x, (0..2_i64.pow(m as u32))))
        .flat_map(|(x, zs)| zs.into_iter().map(move |z| (x, reduce(H(x, z), m), z)));

    //let e = |k: u32, x| (x % 2_i32.pow(k)) as f64 / 2_f64.powf(k as f64);
    chart
        .draw_series(PointSeries::of_element(
            points.map(|(x, y, z)| (x, y, z)),
            1,
            &RED,
            &|c, s, st| Circle::new(c, s, st.filled()),
        ))?
        .label(format!("H(x, y) with m = {}", m))
        .legend(|(x, y)| PathElement::new(vec![(x, y), (x + 20, y)], &RED));

    chart
        .configure_series_labels()
        .background_style(&WHITE.mix(0.8))
        .border_style(&BLACK)
        .draw()?;

    // Hash stuff
    // Document.
    let w = fs::read_to_string("input.txt")
        .unwrap()
        .as_bytes()
        .iter()
        .fold("".to_owned(), |acc, b| {
            format!("{}{}", acc, format!("{:b}", *b))
        })
        .chars()
        .collect::<Vec<_>>();

    let mut h0_bits = vec![];
    let mut rng = rand::thread_rng();
    for _ in 0..m {
        let b = rng.gen_bool(0.5);
        h0_bits.push(if b { '1' } else { '0' })
    }

    let h0 = bin_to_dec(h0_bits);

    //println!("w: {:?}", w.iter().collect::<String>());
    println!("h0: {:0m$b}", h0, m = m);

    let mut hi = h0;
    let a = split(w.clone(), m);
    for ai in a {
        hi = reduce(H(ai, hi), m);
        //println!("hi: {:0m$b}", hi, m = m);
    }

    println!("hk: {:0m$b}", hi, m = m);

    Ok(())
}

fn bin_to_dec(bin: Vec<char>) -> i64 {
    bin.iter().rev().enumerate().fold(0_i64, |acc, (i, c)| {
        acc + c.to_digit(2).unwrap() as i64 * 2_i64.pow(i as u32)
    })
}

/// Split the document `w` on parts of `m` length as vectors of integers.
fn split(w: Vec<char>, m: usize) -> Vec<i64> {
    let mut a: Vec<i64> = vec![];

    let mut rev_w = w.into_iter().rev().collect::<Vec<char>>();

    while !rev_w.is_empty() {
        let len = m.min(rev_w.len());
        let block = rev_w[0..len]
            .to_vec()
            .into_iter()
            .rev()
            .collect::<Vec<char>>();
        a.push(bin_to_dec(block));
        rev_w = rev_w[len..rev_w.len()].to_vec();
        //println!("{:?}", rev_w.iter().collect::<String>());
    }

    a
}

fn reduce(x: i64, m: usize) -> i64 {
    x.rem_euclid(2_i64.pow(m as u32))
}

fn f(x: i64, y: i64) -> i64 {
    sum(1337, x ^ y) * sum(sum_of_squares(x - y, y), exp(x + y, 2))
}

fn sum(x: i64, y: i64) -> i64 {
    x + y
}

fn sum_of_squares(x: i64, y: i64) -> i64 {
    sum(x * x, y * y)
}

fn exp(x: i64, y: i64) -> i64 {
    ((1 + 2 * x) as f64).powi(y as i32) as i64
}

fn div(x: i64, y: i64) -> i64 {
    x * exp(y, -1)
}
