package main

import (
	"bufio"
	"container/heap"
	"flag"
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

var terrainColor = map[rune]string{
	'R': "\033[33m",      // Дорога (жёлтый)
	'G': "\033[32m",      // Поле (зелёный)
	'S': "\033[38;5;94m", // Болото (коричневый)
	'H': "\033[92m",      // Холмы (светло-зелёный)
	'F': "\033[32;1m",    // Лес (тёмно-зелёный)
	'W': "\033[36m",      // Вода (голубой)
	'M': "\033[90m",      // Горы (серый)
	'X': "\033[91m",      // Точка сбора (ярко-красный)
	'B': "\033[97m",      // Белый цвет
	'0': "\033[0m",       // Сброс цвета
}

func main() {
	mazePath := flag.String("m", "maze.txt", "Путь до файла с лабиринтом")
	flag.Parse()

	maze, heroes := readMaze(*mazePath)
	fmt.Println("Лабиринт загружен")
	fmt.Printf("Герои: %v\n", heroes)

	meetingPoint := findOptimalMeetingPoint(maze, heroes)
	fmt.Printf("Оптимальная точка сбора: %+v\n", meetingPoint)

	paths := make([][]Point, len(heroes))
	for i, hero := range heroes {
		paths[i] = findPath(maze, hero, meetingPoint)
		if paths[i] == nil {
			fmt.Printf("Путь для героя %d не найден.\n", i+1)
			continue
		}
		fmt.Printf("Путь героя %d до точки сбора:\n", i+1)
		for _, p := range paths[i] {
			fmt.Printf("x:%d, y:%d\n", p.X, p.Y)
		}
		fmt.Println()
	}

	printMazeWithPaths(maze, heroes, paths)
}

func readMaze(filename string) ([][]rune, []Point) {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Ошибка при открытии файла: %v", err))
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
				terrainColor[cell] = fmt.Sprintf("\033[9%dm", cell-'0')
			}
		}
	}

	return maze, heroes
}

func isValid(p Point, maze [][]rune) bool {
	return p.Y >= 0 && p.Y < len(maze) && p.X >= 0 && p.X < len(maze[0]) && terrainCost[maze[p.Y][p.X]] != math.Inf(1)
}

func findOptimalMeetingPoint(maze [][]rune, heroes []Point) Point {
	minCost := math.Inf(1)
	bestPoint := Point{0, 0}

	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[0]); x++ {
			if terrainCost[maze[y][x]] == math.Inf(1) {
				continue
			}

			totalCost := 0.0
			for _, hero := range heroes {
				path := findPath(maze, hero, Point{x, y})
				if path == nil {
					totalCost = math.Inf(1)
					break
				}
				for _, p := range path {
					totalCost += terrainCost[maze[p.Y][p.X]]
				}
			}

			if totalCost < minCost {
				minCost = totalCost
				bestPoint = Point{x, y}
			}
		}
	}

	return bestPoint
}

type Item struct {
	point Point
	cost  float64
	path  []Point
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i]; pq[i].index, pq[j].index = i, j }
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
func printMazeWithPaths(maze [][]rune, heroes []Point, paths [][]Point) {
	pathMap := make(map[Point]bool)
	for _, path := range paths {
		for _, p := range path {
			pathMap[p] = true
		}
	}

	for y, row := range maze {
		for x, cell := range row {
			p := Point{x, y}
			if pathMap[p] {
				fmt.Print("\033[41m") // Красный фон для пути
			}
			fmt.Print(terrainColor[cell], string(cell), "\033[0m")
		}
		fmt.Println()
	}
}

func findPath(maze [][]rune, start, end Point) []Point {
	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{start, 0, []Point{start}, 0})

	visited := make(map[Point]bool)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)

		if current.point == end {
			return current.path
		}

		if visited[current.point] {
			continue
		}
		visited[current.point] = true

		for i := 0; i < 4; i++ {
			nx, ny := current.point.X+dx[i], current.point.Y+dy[i]
			neighbor := Point{nx, ny}

			if isValid(neighbor, maze) && !visited[neighbor] {
				newCost := current.cost + terrainCost[maze[ny][nx]]
				newPath := append([]Point(nil), current.path...)
				newPath = append(newPath, neighbor)
				heap.Push(pq, &Item{neighbor, newCost, newPath, 0})
			}
		}
	}

	return nil
}
