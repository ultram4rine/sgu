mod funcs;
mod helpers;

use funcs::{f, f_256, g};
use helpers::{adj_from_func_modulo, find_cycle};

fn main() {
    println!("-71 * x - (x * x)");
    if is_bijective(f) {
        println!("bijective")
    } else {
        println!("non bijective")
    }
    if is_transitive(f) {
        println!("transitive")
    } else {
        println!("non transitive")
    }
    println!("");

    println!("77 * x - 7");
    if is_bijective(g) {
        println!("bijective")
    } else {
        println!("non bijective")
    }
    if is_transitive(g) {
        println!("transitive")
    } else {
        println!("non transitive")
    }
    println!("");

    println!(
        "(x ^ 1)
    ^ (2 * (x
        & (1 + 2 * x)
        & (3 + 4 * x)
        & (7 + 8 * x)
        & (15 + 16 * x)
        & (31 + 32 * x)
        & (63 + 64 * x)))
    ^ (4 * (x * x + 27))"
    );
    if is_transitive_modulo(f_256, 256) {
        println!("f(x) transitive with modulo 256")
    } else {
        println!("f(x) non transitive with modulo 256")
    }
}

fn is_bijective(f: fn(i32) -> i32) -> bool {
    is_bijective_modulo(f, 4)
}

fn is_transitive(f: fn(i32) -> i32) -> bool {
    is_transitive_modulo(f, 8)
}

fn is_bijective_modulo(f: fn(i32) -> i32, modulo: i32) -> bool {
    let adj: Vec<i32> = adj_from_func_modulo(f, modulo, false);

    let mut cycles_from = vec![false; modulo as usize];
    let mut colors = vec!["white".to_owned(); modulo as usize];
    for i in 0..modulo {
        if colors[i as usize] == "white" {
            if find_cycle(i, adj.clone(), &mut colors) {
                for j in 0..modulo {
                    if colors[j as usize] == "black".to_owned() {
                        cycles_from[j as usize] = true;
                    }
                }
            } else {
                continue;
            };
        }
    }

    let cycles_everywhere = cycles_from.iter().fold(true, |res, x| res && *x);

    if cycles_everywhere {
        true
    } else {
        false
    }
}

fn is_transitive_modulo(f: fn(i32) -> i32, modulo: i32) -> bool {
    let adj: Vec<i32> = adj_from_func_modulo(f, modulo, false);

    let mut colors = vec!["white".to_owned(); modulo as usize];

    // If after DFS all nodes painted to black, then graph is cycle itself.
    if find_cycle(0, adj.clone(), &mut colors) {
        let circular = colors
            .iter()
            .fold(true, |res, x| res && (*x == "black".to_owned()));
        if circular {
            true
        } else {
            false
        }
    } else {
        false
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // Test funcs

    /// Bijective.
    pub fn bijective(x: i32) -> i32 {
        x + 6
    }

    /// Transitive.
    pub fn transitive(x: i32) -> i32 {
        x + 1
    }

    /// Not biective, not transitive.
    pub fn non(x: i32) -> i32 {
        x * x + 1
    }

    #[test]
    fn test_is_bijective() {
        assert_eq!(is_bijective(bijective), true);
    }

    #[test]
    fn test_is_transitive() {
        assert_eq!(is_transitive(transitive), true);
    }

    #[test]
    fn test_is_not_bijective() {
        assert_eq!(is_bijective(non), false);
    }

    #[test]
    fn test_is_not_transitive() {
        assert_eq!(is_transitive(non), false);
    }
}
