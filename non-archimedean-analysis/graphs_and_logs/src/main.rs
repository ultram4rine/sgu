#![allow(non_snake_case)]
use std::fs::File;

fn main() {
    let p: usize = 11;
    let n: usize = p - 1;
    let a: usize = 2;

    let X = log(p, n, a);
    println!("X: {}", X);

    let mut nodes = vec![];
    let mut colors = vec![];
    let mut edges = vec![];

    for x in 0..2_u64.pow(n as u32) {
        nodes.push(x as usize);
        if x == X as u64 {
            colors.push("red");
        } else {
            colors.push("black");
        }
        edges.push((x as usize, A(x, n) as usize));
    }

    let graph = Graph {
        nodes,
        colors,
        edges,
    };

    let mut file = File::create("graph.dot").unwrap();
    dot::render(&graph, &mut file).unwrap()
}

fn A(x: u64, n: usize) -> u64 {
    // TODO: actually don't need to extend anymore.
    let bin_x_tmp = format!("{:b}", x)
        .to_owned()
        .chars()
        .map(|c| c.to_digit(2).unwrap() as u8)
        .collect::<Vec<u8>>();

    let bin_x = if bin_x_tmp.len() < n {
        let diff = n - bin_x_tmp.len();
        let mut tmp = vec![0 as u8; diff];
        tmp.extend(bin_x_tmp);
        tmp
    } else {
        bin_x_tmp
    };

    let mut bin_y = vec![];
    for j in 0..n {
        let c = (bin_x[(j + 1) % n] ^ bin_x[j]) as u8;
        bin_y.push(c);
    }

    let y = bin_y.iter().enumerate().fold(0, |acc, (i, xi)| {
        acc + *xi as i64 * 2_i64.pow((n - i - 1) as u32)
    });

    //println!("{:?} -> {}", bin_y, y);

    y as u64
}

fn log(p: usize, n: usize, a: usize) -> usize {
    let mut set = vec![];
    for i in 1..p {
        set.push(a.pow(i as u32) % p)
    }
    println!("{:?}", set);
    set.sort();

    let mut ls = vec![];
    'k: for k in set {
        for i in 1..=n {
            if a.pow(i as u32) % p == k {
                ls.push(i);
                continue 'k;
            }
        }
    }

    println!("{:?}", ls);

    ls.iter().enumerate().fold(0, |acc, (i, l)| {
        acc + (*l % 2) * 2_usize.pow((n - i - 1) as u32)
    })
}

// Needed for graphviz stuff.
type Node = usize;
type Edge<'a> = &'a (Node, Node);
struct Graph<'a> {
    nodes: Vec<usize>,
    colors: Vec<&'a str>,
    edges: Vec<(usize, usize)>,
}

impl<'a> dot::Labeller<'a, Node, Edge<'a>> for Graph<'a> {
    fn graph_id(&'a self) -> dot::Id<'a> {
        dot::Id::new("result").unwrap()
    }
    fn node_id(&'a self, n: &Node) -> dot::Id<'a> {
        dot::Id::new(format!("N{}", n)).unwrap()
    }
    fn node_color(&'a self, n: &Node) -> Option<dot::LabelText<'a>> {
        Some(dot::LabelText::LabelStr(std::borrow::Cow::Borrowed(
            self.colors[*n],
        )))
    }
    fn node_label<'b>(&'b self, n: &Node) -> dot::LabelText<'b> {
        dot::LabelText::LabelStr(self.nodes[*n].to_string().into())
    }
    fn edge_label<'b>(&'b self, _: &Edge) -> dot::LabelText<'b> {
        dot::LabelText::LabelStr("".into())
    }
}

impl<'a> dot::GraphWalk<'a, Node, Edge<'a>> for Graph<'a> {
    fn nodes(&self) -> dot::Nodes<'a, Node> {
        (0..self.nodes.len()).collect()
    }
    fn edges(&'a self) -> dot::Edges<'a, Edge<'a>> {
        self.edges.iter().collect()
    }
    fn source(&self, e: &Edge) -> Node {
        e.0
    }
    fn target(&self, e: &Edge) -> Node {
        e.1
    }
}
