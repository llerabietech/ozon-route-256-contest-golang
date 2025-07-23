package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Структура для хранения координат робота
// x — номер строки, y — номер столбца
// (0,0) — верхний левый угол, (n-1,m-1) — нижний правый угол

type Position struct {
	x int
	y int
}

// Основная функция обработки входных данных и построения маршрутов
func process(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscanln(in, &n)

	for i := 0; i < n; i++ {
		// Инициализация позиций роботов
		a := Position{}
		b := Position{}

		var xs, ys int
		fmt.Fscanln(in, &xs, &ys)

		// Чтение сетки склада
		grid := make([][]string, xs)
		for i := range grid {
			grid[i] = make([]string, ys)
		}

		var line string
		for x := 0; x < xs; x++ {
			fmt.Fscanln(in, &line)
			for y := 0; y < ys; y++ {
				grid[x][y] = string(line[y])
				if grid[x][y] == "A" {
					a.x = x
					a.y = y
				}
				if grid[x][y] == "B" {
					b.x = x
					b.y = y
				}
			}
		}

		// В зависимости от положения роботов вызываем соответствующие функции построения маршрута
		if a.x == 0 && a.y == 0 {
			// Если A уже в левом верхнем углу, строим путь для B к правому нижнему
			goDownRight(grid, b, "b")
		} else if a.x == len(grid)-1 && a.y == len(grid[0])-1 {
			// Если A уже в правом нижнем углу, строим путь для B к левому верхнему
			goUpLeft(grid, b, "b")
		} else if b.x == 0 && b.y == 0 {
			// Если B уже в левом верхнем углу, строим путь для A к правому нижнему
			goDownRight(grid, a, "a")
		} else if b.x == len(grid)-1 && b.y == len(grid[0])-1 {
			// Если B уже в правом нижнем углу, строим путь для A к левому верхнему
			goUpLeft(grid, a, "a")
		} else if a.x == b.x {
			// Если роботы на одной строке, сначала строим путь для того, кто левее
			if a.y < b.y {
				goUpLeft(grid, a, "a")
				goDownRight(grid, b, "b")
			} else {
				goUpLeft(grid, b, "b")
				goDownRight(grid, a, "a")
			}
		} else if a.x <= b.x {
			// Если A выше B, сначала строим путь для A к левому верхнему, затем для B к правому нижнему
			goUpLeft(grid, a, "a")
			goDownRight(grid, b, "b")
		} else {
			// В остальных случаях сначала строим путь для B к левому верхнему, затем для A к правому нижнему
			goUpLeft(grid, b, "b")
			goDownRight(grid, a, "a")
		}

		// Вывод результата для текущего теста
		out.WriteString(gridToString(grid))
	}
}

// Функция для построения пути к левому верхнему углу
func goUpLeft(grid [][]string, pos Position, mark string) {
	// Если сверху препятствие, сначала двигаемся влево
	if pos.x > 0 {
		if grid[pos.x-1][pos.y] == "#" {
			pos.y--
			grid[pos.x][pos.y] = mark
		}
	}
	pos = goUp(grid, pos, mark)
	pos = goLeft(grid, pos, mark)
}

// Движение вверх до границы
func goUp(grid [][]string, pos Position, mark string) Position {
	for i := pos.x - 1; i >= 0; i-- {
		pos.x = i
		grid[pos.x][pos.y] = mark
	}
	return pos
}

// Движение влево до границы
func goLeft(grid [][]string, pos Position, mark string) Position {
	for i := pos.y - 1; i >= 0; i-- {
		pos.y = i
		grid[pos.x][pos.y] = mark
	}
	return pos
}

// Функция для построения пути к правому нижнему углу
func goDownRight(grid [][]string, pos Position, mark string) {
	// Если снизу препятствие, сначала двигаемся вправо
	if pos.x < len(grid)-2 {
		if grid[pos.x+1][pos.y] == "#" {
			pos.y++
			grid[pos.x][pos.y] = mark
		}
	}
	pos = goDown(grid, pos, mark)
	pos = goRight(grid, pos, mark)
}

// Движение вниз до границы
func goDown(grid [][]string, pos Position, mark string) Position {
	for i := pos.x + 1; i < len(grid); i++ {
		pos.x = i
		grid[pos.x][pos.y] = mark
	}
	return pos
}

// Движение вправо до границы
func goRight(grid [][]string, pos Position, mark string) Position {
	for i := pos.y + 1; i < len(grid[0]); i++ {
		pos.y = i
		grid[pos.x][pos.y] = mark
	}
	return pos
}

// Преобразование сетки в строку для вывода
func gridToString(grid [][]string) string {
	var sb strings.Builder
	for _, row := range grid {
		sb.WriteString(strings.Join(row, ""))
		sb.WriteString("\n")
	}
	return sb.String()
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	process(in, out)
}
