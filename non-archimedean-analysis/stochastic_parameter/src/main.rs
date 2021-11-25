#[allow(non_snake_case)]

fn main() {
    let n: i64 = 31;
    let a: i64 = 4;

    if !is_coprime(a, n) {
        panic!("a and n should be coprime")
    }

    let gamma = make_gamma(n);
    println!("Gamma({}): {:?}", n, gamma);

    let mut orbits: Vec<Vec<i64>> = vec![];
    let mut used = vec![false; n as usize];
    for x in gamma.clone() {
        if !used[x as usize] {
            used[x as usize] = true;
            let mut orbit: Vec<i64> = vec![];
            orbit.push(x);
            let mut next = f(a, x, n);
            while next != x {
                used[next as usize] = true;
                orbit.push(next);
                next = f(a, next, n)
            }
            orbit.sort();
            orbits.push(orbit);
        }
    }

    let mut target_orbit = vec![];
    let mut l = 0;
    let first = target_orbit_evolution(a, l, n);
    target_orbit.push(first);
    l += 1;
    let mut next = target_orbit_evolution(a, l, n);
    while next != first {
        l += 1;
        target_orbit.push(next);
        next = target_orbit_evolution(a, l, n);
    }
    target_orbit.sort();
    println!("Target: {:?}", target_orbit);

    let T = orbits.first().unwrap().len();
    println!("T({}): {}\n", n, T);

    let m = gamma.len();
    for orbit in orbits {
        if orbit == target_orbit {
            println!("Target orbit: {:?}", orbit);
        } else {
            println!("Orbit: {:?}", orbit);
        }

        let pos_first = gamma
            .iter()
            .position(|&g| g == *orbit.first().unwrap())
            .unwrap()
            + 1;
        let pos_last = gamma
            .iter()
            .position(|&g| g == *orbit.last().unwrap())
            .unwrap()
            + 1;
        let mut R = (m - pos_last + pos_first).pow(2);
        for i in 0..orbit.len() - 1 {
            let pos = gamma.iter().position(|&g| g == orbit[i]).unwrap() + 1;
            let pos_next = gamma.iter().position(|&g| g == orbit[i + 1]).unwrap() + 1;
            R += (pos_next - pos).pow(2);
        }
        println!("R: {}", R);

        let r = R as f64 / (m * m) as f64;
        println!("r: {}", r);

        let s = r * T as f64;
        println!("s: {}\n", s);
    }
}

fn target_orbit_evolution(a: i64, l: i64, n: i64) -> i64 {
    a.pow(l as u32) % n
}

/// Evolution function.
fn f(a: i64, x: i64, n: i64) -> i64 {
    (a * x) % n
}

fn make_gamma(n: i64) -> Vec<i64> {
    let mut gamma: Vec<i64> = vec![];

    for i in 1..n {
        if is_coprime(i, n) {
            gamma.push(i)
        }
    }

    gamma
}

fn is_coprime(a: i64, b: i64) -> bool {
    gcd(a, b) == 1
}

fn gcd(a: i64, b: i64) -> i64 {
    if a == 0 || b == 0 {
        0
    } else if a == b {
        a
    } else if a > b {
        gcd(a - b, b)
    } else {
        gcd(a, b - a)
    }
}
