# WikiRace Solver

Finding the Shortest Path from a Wikipedia Article to Another Wikipedia Article with BFS (Breadth-First Search) or IDS (Iterative Deepening Search) algorithm

## BFS (Breadth-First Search)
Breadth-first search (BFS) explores a tree data structure starting from the root, traversing all nodes at the current depth before moving to nodes at the next depth level.  
We implement this algorithm in src/backend/utils/bfs.go

## IDS (Iterative Deepening Search)
Iterative deepening search is a searching strategy where we keep running a limited version of depth-first search (DFS), gradually increasing the depth, until we find the goal.  
We implement this algorithm in src/backend/utils/ids.go

## Requirements

- [Go](https://go.dev/dl/)  
  We use go version 1.22.2.  
  Install this version if you find problems in different version

- A web browser (Mozilla Firefox, Google Chrome, etc.)

## How to Set Up and Run

1. Clone this repository
   ```
   git clone https://github.com/IrfanSidiq/Tubes2_wikiwikiwiki.git
   ```

2. Open terminal in the root project folder (the default is Tubes2_wikiwikiwiki)

3. Change directory to src/backend
   ```
   cd src/backend
   ```

4. There are 2 options:  
   1. Run main.go file directly
      ```
      go run main.go
      ```
   2. Build the go files
      ```
      go build
      ```
      Then, run the executable file
      ```
      ./backend
      ```

5. If there is a firewall security alert, you can ignore it or choose an option that you feel safe to allow

6. Open http://localhost:8080/ in your web browser

## Authors

Group name: wikiwikiwiki

1. [Irfan Sidiq Permana](https://github.com/IrfanSidiq)  
   13522007

2. [Bastian H. Suryapratama](https://github.com/bastianhs)  
   13522034

3. [Diero Arga Purnama](https://github.com/DieroA)  
   13522056
