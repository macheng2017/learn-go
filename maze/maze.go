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

// 定义出点的结构体,不要用x,y因为我们定义的坐标轴正好和笛卡尔坐标系相反,使用x,y容易混淆
type point struct {
	i, j int
}

// 定义一个函数,将当前的坐标值与上左下右的坐标相加
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// 定义出上左下右这四个方向
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) at(grid [][]int) (int, bool) {
	// 使用bool表示是否越界(地图上是否越界)
	if p.i < 0 || p.j >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

// 将地图传入行走函数,并且定义出起点和结束位置
func walk(maze [][]int, start, end point) {
	//定义一个 step二维数组,用来存放步数,这个概念很重要,最终的路径都是有一个个步数连成的
	// 定义方式和上面的地图中的二维数组方法一样
	steps := make([][]int, len(maze))
	for i := range steps {
		// 初始化后的二维数组都是0,这里和地图中的带围墙(数字1)不太一样
		steps[i] = make([]int, len(maze[i]))
	}
	// 定义一个队列,将起点加入到队列当中作为初始值
	Q := []point{start}
	// 定义一个循环并写好退出条件
	// 1. 队列为空 2. 到达终点
	for len(Q) > 0 {
		// 当前节点
		cur := Q[0]
		Q = Q[1:]
		// 按照上左下右的顺序探索地图,先定义出这四个方向,然后相加(发现新的节点)

		for _, dir := range dirs {
			next := cur.add(dir)
			// maze at next is 0
			// and steps at next is 0
			// and next != start
			// 必须满足上面的条件才能走过去,这样条件比较多
			// 可以反过来,把不满足条件的剔除

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

		}

	}

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
