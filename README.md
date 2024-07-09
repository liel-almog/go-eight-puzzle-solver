# Eight Puzzle Solver

This project is a web-based solver for the Eight Puzzle problem, implemented in Golang and Typescript. It supports three different algorithms: Breadth-First Search (BFS), Depth-First Search (DFS), and A* Search (ASTAR). The solver can handle puzzles of different sizes.

## Table of Contents

- [Eight Puzzle Solver](#eight-puzzle-solver)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Features](#features)
    - [Web Interface](#web-interface)
  - [Algorithms](#algorithms)
    - [Breadth-First Search (BFS)](#breadth-first-search-bfs)
    - [Depth-First Search (DFS)](#depth-first-search-dfs)
    - [A\* Search (ASTAR)](#a-search-astar)
  - [Usage](#usage)
    - [Get Started with Docker 🐳](#get-started-with-docker-)
  - [Contributing](#contributing)
  - [License](#license)
  - [Acknowledgements](#acknowledgements)

## Introduction

The Eight Puzzle is a sliding puzzle that consists of a grid with numbered tiles and one empty space. The goal is to move the tiles around until they are in numerical order. This project provides a web-based solver for the Eight Puzzle using three different algorithms.

## Features

- Solve the Eight Puzzle using BFS, DFS, and A* Search algorithms.
- Support for different puzzle sizes.
- Web-based interface for easy usage.
- Ability to interactively solve the puzzle or use the algorithm to find the solution.
- Implemented in Golang and Typescript.


### Web Interface

1. Select the puzzle size.
2. Create a puzzle configuration - we can change it using the UI. 
3. Choose the algorithm (BFS, DFS, ASTAR).
4. Click the solve button to find a solution using the chosen algorithm.
5. You can interactively move the tiles and try to solve the puzzle yourself or let the algorithm solve it.
6. The solution path and steps will be displayed on the interface.

Here is a screenshot of the web interface:

<img src="https://github.com/liel-almog/go-eight-puzzle-solver/blob/main/public/initial_state.png" width="400" height="400">
<img src="https://github.com/liel-almog/go-eight-puzzle-solver/blob/main/public/solved_state.png" width="400" height="400">


## Algorithms

### Breadth-First Search (BFS)

Breadth-First Search (BFS) is an algorithm for traversing or searching tree or graph data structures. It starts at the root node and explores all neighbor nodes at the present depth level before moving on to nodes at the next depth level. It returns the shortest path to the target.

### Depth-First Search (DFS)

Depth-First Search (DFS) is an algorithm for traversing or searching tree or graph data structures. It starts at the root node and explores as far as possible along each branch before backtracking.

### A* Search (ASTAR)

A* Search is a computer algorithm that is widely used in pathfinding and graph traversal. It is an extension of Dijkstra's algorithm that uses heuristics to improve performance. It is used the find the **easiest path - shortest path (with respect to the given weights)** from a source to the target

## Usage

### Get Started with Docker 🐳

Getting the eight puzzle solver up and running is as simple as executing a single command, thanks to Docker Compose.

1. Clone the repository
   ```
   git clone https://github.com/liel-almog/go-eight-puzzle-solver.git
   ```
2. Navigate to the project directory
   ```
   cd go-eight-puzzle-solver
   ```
3. Start the application
   ```
   docker compose up
   ```

Visit `http://localhost:3000` to access the eight puzzle solver web interface.


## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you would like to contribute.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Golang](https://golang.org/)
- [Node.js](https://nodejs.org/)
- [npm](https://www.npmjs.com/)
