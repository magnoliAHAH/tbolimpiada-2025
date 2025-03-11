package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

func main() {
	maze, start, exit := readMaze("maze.txt")

	path := bfs(maze, start, exit)
	if path != nil {
		printPath(path)
	} else {
		fmt.Println("Путь не найден.")
	}
}

func readMaze(filename string) ([][]rune, Point, Point) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var width, height int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &width)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &height)

	maze := make([][]rune, height)
	var start, exit Point

	for y := 0; y < height; y++ {
		scanner.Scan()
		line := scanner.Text()
		maze[y] = []rune(line)
		for x, cell := range line {
			switch cell {
			case '1':
				start = Point{x, y}
			case 'F':
				exit = Point{x, y}
			}
		}
	}

	return maze, start, exit
}

func bfs(maze [][]rune, start, exit Point) []Point {
	directions := []Point{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	queue := []Point{start}
	visited := make(map[Point]bool)
	visited[start] = true
	prev := make(map[Point]Point)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == exit {
			return reconstructPath(prev, start, exit)
		}

		for _, dir := range directions {
			next := Point{current.X + dir.X, current.Y + dir.Y}

			if isValid(next, maze, visited) {
				visited[next] = true
				queue = append(queue, next)
				prev[next] = current
			}
		}
	}

	return nil
}

func isValid(p Point, maze [][]rune, visited map[Point]bool) bool {
	if p.Y < 0 || p.Y >= len(maze) || p.X < 0 || p.X >= len(maze[0]) {
		return false
	}
	if maze[p.Y][p.X] == '#' || visited[p] {
		return false
	}
	return true
}

func reconstructPath(prev map[Point]Point, start, exit Point) []Point {
	var path []Point
	for at := exit; at != start; at = prev[at] {
		path = append(path, at)
	}
	path = append(path, start)

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func printPath(path []Point) {
	for _, p := range path {
		fmt.Printf("x:%d, y:%d\n", p.X+1, p.Y+1)
	}
}
