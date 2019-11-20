package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, e := os.Open(filename)

	if e != nil {
		// 在这里不做错误处理了,假定文件名字和内容都是正确的
		panic(e)
	}
	var row, col int
	// 按照给定的格式读出文件中的内容
	fmt.Fscanf(file, "%d %d", &row, &col)
	// 首先读出来的是初始化二维数组的行和列

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			//使用两层for循环确定一个值的位置
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

// 定义出点的结构体
type point struct {
	i, j int
}

// 将地图传入行走函数,并且定义出起点和结束位置
func walk(maze [][]int, start, end point) {

}

func main() {
	res := readMaze("maze/maze.in")
	// 起始位置是地图的左上角,定义出点的结构体
	start := point{0, 0}
	end := point{len(res) - 1, len(res[0]) - 1}
	walk(res, start, end)

}

//0 1 0 0 0
//0 0 0 1 0
//0 1 0 1 0
//1 1 1 0 0
//0 1 0 0 1
//0 1 0 0 0
