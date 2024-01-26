// Implementing Dijkstra's algorithm in Rust for adjacency list

use std::cmp::Ordering;
use std::usize::MAX;
use std::collections::BinaryHeap;

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

fn dijkstra(graph: Vec<Vec<(usize, usize)>>, start: usize, end: usize) -> Vec<usize> {
    let n = graph.len();
    let mut dist = vec![MAX; n];
    let mut visited = vec![false; n];
    let mut min_heap = BinaryHeap::new();

    dist[start] = 0;
    min_heap.push(Node { index: start, distance: 0 });

    while let Some(Node { index, distance }) = min_heap.pop() {
        if visited[index] {
            continue;
        }

        visited[index] = true;

        for &(neighbor, weight) in &graph[index] {
            if dist[index] + weight < dist[neighbor] {
                dist[neighbor] = dist[index] + weight;
                min_heap.push(Node { index: neighbor, distance: dist[neighbor] });
            }
        }
    }

    dist
}

fn main() {
    let graph = vec![
        vec![(1, 7), (2, 5), (3, 2), (6, 2), (7, 4), (8, 3)],
        vec![(0, 7), (4, 3)],
        vec![(0, 5), (3, 10), (4, 4)],
        vec![(0, 2), (2, 10), (5, 2)],
        vec![(1, 3), (2, 4), (5, 6)],
        vec![(1, 8), (3, 2), (4, 6)],
        vec![(0, 2), (7, 5)],
        vec![(6, 5), (8, 1)],
        vec![(7, 1)],
    ];

    println!("The given nodes are: {:?}", graph);
    let start = 2;
    let end = 3;

    let dist = dijkstra(graph, start, end);

    println!("Shortest path from node {} to {}: {}", start, end, dist[end]);
}
