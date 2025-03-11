package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	X, Y int
}

var terrainCost = map[rune]float64{
	'R': 0.5,         // Дорога
	'G': 1,           // Поле
	'S': 5,           // Болото
	'H': 4,           // Холмы
	'F': 3,           // Лес
	'W': math.Inf(1), // Вода (непроходимо)
	'M': math.Inf(1), // Горы (непроходимо)
}

func main() {
	maze, heroes := readMaze("maze.txt")
	fmt.Println("Лабиринт загружен")
	fmt.Printf("Герои: %v\n", heroes)

	meetingPoint := findOptimalMeetingPoint(maze, heroes)
	fmt.Printf("Оптимальная точка сбора: %+v\n", meetingPoint)

	for i, hero := range heroes {
		path := findPath(maze, hero, meetingPoint)
		fmt.Printf("Путь героя %d до точки сбора:\n", i+1)
		for _, p := range path {
			fmt.Printf("x:%d, y:%d\n", p.X, p.Y)
		}
		fmt.Println()
	}

	printMazeWithPaths(maze, heroes)
}

func readMaze(filename string) ([][]rune, []Point) {
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
	var heroes []Point

	for y := 0; y < height; y++ {
		scanner.Scan()
		line := scanner.Text()
		maze[y] = []rune(line)
		for x, cell := range line {
			if cell >= '1' && cell <= '9' {
				heroes = append(heroes, Point{x, y})
			}
		}
	}

	return maze, heroes
}

func findOptimalMeetingPoint(maze [][]rune, heroes []Point) Point {
	height, width := len(maze), len(maze[0])
	bestPoint := Point{0, 0}
	minTotalCost := math.Inf(1)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if !isValid(Point{x, y}, maze) {
				continue
			}

			totalCost := 0.0
			valid := true
			for _, hero := range heroes {
				distances := bfs(maze, hero)
				cost := distances[y][x]
				if math.IsInf(cost, 1) {
					valid = false
					break
				}
				totalCost += cost
			}

			if valid && totalCost < minTotalCost {
				minTotalCost = totalCost
				bestPoint = Point{x, y}
			}
		}
	}

	return bestPoint
}

func findPath(maze [][]rune, start, target Point) []Point {
	path := []Point{}
	distances := bfs(maze, start)
	current := target

	if math.IsInf(distances[current.Y][current.X], 1) {
		return path
	}
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for current != start {
		path = append([]Point{current}, path...)
		bestNext := current
		bestCost := distances[current.Y][current.X]

		for _, dir := range directions {
			next := Point{current.X + dir.X, current.Y + dir.Y}
			if isValid(next, maze) && distances[next.Y][next.X] < bestCost {
				bestNext = next
				bestCost = distances[next.Y][next.X]
			}
		}

		if bestNext == current {
			break
		}
		current = bestNext
	}

	path = append([]Point{start}, path...)
	return path
}

func bfs(maze [][]rune, start Point) [][]float64 {
	height, width := len(maze), len(maze[0])
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	distances := make([][]float64, height)
	for y := range distances {
		distances[y] = make([]float64, width)
		for x := range distances[y] {
			distances[y][x] = math.Inf(1)
		}
	}
	distances[start.Y][start.X] = 0

	queue := []Point{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			next := Point{current.X + dir.X, current.Y + dir.Y}

			if isValid(next, maze) {
				newCost := distances[current.Y][current.X] + terrainCost[maze[next.Y][next.X]]
				if newCost < distances[next.Y][next.X] {
					distances[next.Y][next.X] = newCost
					queue = append(queue, next)
				}
			}
		}
	}

	return distances
}

func isValid(p Point, maze [][]rune) bool {
	return p.Y >= 0 && p.Y < len(maze) && p.X >= 0 && p.X < len(maze[0]) && terrainCost[maze[p.Y][p.X]] != math.Inf(1)
}

func printMazeWithPaths(maze [][]rune, heroes []Point) {
	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			cell := maze[y][x]
			for i, hero := range heroes {
				if hero.X == x && hero.Y == y {
					cell = rune('1' + i)
				}
			}
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}
