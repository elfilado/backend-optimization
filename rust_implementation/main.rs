// Implementing Dijkstra's algorithm in Rust for adjacency matrix

use std::cmp::Ordering;
use std::usize::MAX;

#[derive(Debug, Copy, Clone, Eq, PartialEq)]
struct Node {
    index: usize,
    distance: usize,
}

impl Ord for Node {
    fn cmp(&self, other: &Self) -> Ordering {
        other.distance.cmp(&self.distance)
    }
}

impl PartialOrd for Node {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

fn dijkstra(graph: Vec<Vec<usize>>, start: usize, end: usize) -> Vec<usize> {
    let n = graph.len();
    let mut dist = vec![MAX; n];
    let mut visited = vec![false; n];

    dist[start] = 0;

    for _ in 0..n - 1 {
        let mut u = n;
        for (i, &d) in dist.iter().enumerate() {
            if !visited[i] && (u == n || d < dist[u]) {
                u = i;
            }
        }

        if u == n {
            break;
        }

        visited[u] = true;

        for v in 0..n {
            if graph[u][v] != 0 && dist[u] + graph[u][v] < dist[v] {
                dist[v] = dist[u] + graph[u][v];
            }
        }
    }

    dist
}

fn main() {
    let graph = vec![
        vec![0, 7, 5, 2, 0, 0],
        vec![7, 0, 0, 0, 3, 0],
        vec![5, 0, 0, 10, 4, 0],
        vec![2, 0, 10, 0, 0, 2],
        vec![0, 3, 4, 0, 0, 6],
        vec![0, 8, 0, 2, 6, 0],
    ];
    println!("The given nodes are: {:?}", graph);
    let start = 2;
    let end = 3;

    let dist = dijkstra(graph, start, end);

    println!("Shortest path from node {} to {}: {}", start, end, dist[end]);
}
