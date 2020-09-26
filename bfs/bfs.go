package main

import (
	"fmt"
	"os"
	"strconv"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	var end string
	// // 读取一行，遇到换行符停止读取; end 把回车读取; 遇到空格跳过，遇到换行停止，就是取一行数据
	fmt.Fscanf(file, "%d %d %s", &row, &col, &end)
	maze := make([][]int, row) //第一个 [] 是一个 slice ，   slice的元素类型是  []int
	for i := range maze {      //针对每一行作出处理
		maze[i] = make([]int, col) //表示每一行都有col列
		for j := range maze[i] {   //获得每一行的列数
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
		fmt.Fscanf(file, "%d", &end)
	}

	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}
func (p point) printPath(path string,maze [][] int) string {
	if p.i == len(maze) - 1 && p.j == len(maze[0]) - 1 {
		path += "["  + strconv.Itoa(p.i) + "," + strconv.Itoa(p.j)+ "]"
	}else {
		path += "["  + strconv.Itoa(p.i) + "," + strconv.Itoa(p.j)+ "] -> "
	}
	return path
}
func walk(maze [][]int,
	start, end point) ([][]int,string) {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}
	var path string
	path = start.printPath(path,maze)
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] =
				curSteps + 1
			path = next.printPath(path,maze)
			Q = append(Q, next)
		}
	}

	return steps,path
}

func main() {
	maze := readMaze("bfs/maze/maze.in")

	steps,path := walk(maze, point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%5d", val)
		}
		fmt.Println()
	}
	fmt.Println(path)
	// TODO: construct path from steps
}
