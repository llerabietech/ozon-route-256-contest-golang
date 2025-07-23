package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Process(in, out)
}

func Process(in *bufio.Reader, out *bufio.Writer) {
	var t int
	fmt.Fscanln(in, &t)
	for test := 0; test < t; test++ {
		var n, m int
		fmt.Fscanln(in, &n, &m)
		field := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Fscanln(in, &field[i])
		}

		// Сохраняем координаты всех гексагонов по регионам
		type coord struct{ x, y int }
		regions := make(map[byte][]coord)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				c := field[i][j]
				if c >= 'A' && c <= 'Z' {
					regions[c] = append(regions[c], coord{i, j})
				}
			}
		}

		// Смещения для соседей (гексагональная сетка)
		// Для чётных и нечётных строк разный набор соседей
		var evenD = [][2]int{{-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, 0}, {1, 1}}  // чётная строка (i%2==0)
		var oddD = [][2]int{{-1, -1}, {-1, 0}, {0, -1}, {0, 1}, {1, -1}, {1, 0}} // нечётная строка (i%2==1)

		ok := true
		used := make([][]bool, n)
		for i := range used {
			used[i] = make([]bool, m)
		}

		for region, cells := range regions {
			// Сбросить used для этого региона
			for i := range used {
				for j := range used[i] {
					used[i][j] = false
				}
			}
			// BFS
			queue := []coord{cells[0]}
			used[cells[0].x][cells[0].y] = true
			count := 1
			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]
				var dirs [][2]int
				if cur.x%2 == 1 {
					dirs = evenD
				} else {
					dirs = oddD
				}
				for _, d := range dirs {
					nx, ny := cur.x+d[0], cur.y+d[1]
					if nx >= 0 && nx < n && ny >= 0 && ny < m && !used[nx][ny] && field[nx][ny] == region {
						used[nx][ny] = true
						queue = append(queue, coord{nx, ny})
						count++
					}
				}
			}
			if count != len(cells) {
				ok = false
				break
			}
		}
		if ok {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
