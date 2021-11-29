/// Finds cycle in graph.
///
/// # Arguments
///
/// * `v` - next node of graph.
/// * `adj` - adjacency list.
/// * `colors` - vector(array) of used nodes.
pub fn find_cycle(v: i32, adj: Vec<i32>, mut colors: &mut Vec<String>) -> bool {
    colors[v as usize] = "grey".to_owned();

    let u = adj[v as usize];

    if colors[u as usize] == "white".to_owned() {
        if find_cycle(u, adj.clone(), &mut colors) {
            colors[v as usize] = "black".to_owned();
            return true;
        }
    } else if colors[u as usize] == "grey".to_owned() {
        colors[v as usize] = "black".to_owned();
        return true;
    }

    colors[v as usize] = "black".to_owned();
    false
}

/// Creates graph adjacency list from function and modulo.
///
/// # Arguments
///
/// * `f` - function.
/// * `modulo` - modulo.
/// * `debug` - print adjacency list.
pub fn adj_from_func_modulo(f: fn(i32) -> i32, modulo: i32, debug: bool) -> Vec<i32> {
    let mut adj: Vec<i32> = vec![];
    for k in 0..modulo {
        adj.push(f(k).rem_euclid(modulo));
        if debug {
            println!("{} -> {}", k, adj[k as usize]);
        }
    }
    adj
}
