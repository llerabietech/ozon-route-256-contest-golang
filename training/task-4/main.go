package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Структура для хранения позиции гексагона
type Position struct {
	x int
	y int
}

func Process(in *bufio.Reader, out *bufio.Writer) {
	var t int
	fmt.Fscanln(in, &t)
	for i := 0; i < t; i++ {
		var r, c int
		fmt.Fscanln(in, &r, &c)

		// Читаем карту как двумерный срез строк
		field := make([][]string, 0, r)
		for j := 0; j < r; j++ {
			xStr, _ := in.ReadString('\n')
			xStr = strings.Trim(xStr, "\n")
			x := strings.Split(xStr, "")
			field = append(field, x)
		}

		isValid := checkField(field)
		if isValid {
			fmt.Fprintln(out, "YES")
			continue
		}
		fmt.Fprintln(out, "NO")
	}
}

// Проверяет, что все регионы на карте связны
func checkField(field [][]string) bool {
	hexes := map[string]map[Position]bool{}
	// Собираем все гексагоны по регионам (буквам)
	for r, x := range field {
		for c, v := range x {
			if v != "." {
				p := Position{r, c}
				_, exists := hexes[v]
				if !exists {
					hexes[v] = map[Position]bool{}
				}
				hexes[v][p] = false // изначально не посещён
			}
		}
	}

	// Для каждого региона запускаем обход из одной клетки
	for _, Positiones := range hexes {
		for p := range Positiones {
			visit(Positiones, p)
			break
		}
	}

	// Проверяем, что все клетки каждого региона были посещены
	for _, Positiones := range hexes {
		for _, visited := range Positiones {
			if !visited {
				return false
			}
		}
	}

	return true
}

// Возвращает список соседей для гексагона в гекс-сетке
func getNextMove(p Position) []Position {
	return []Position{
		{p.x, p.y - 2},     // Влево
		{p.x, p.y + 2},     // Вправо
		{p.x - 1, p.y - 1}, // Вверх-влево
		{p.x - 1, p.y + 1}, // Вверх-вправо
		{p.x + 1, p.y - 1}, // Вниз-влево
		{p.x + 1, p.y + 1}, // Вниз-вправо
	}
}

// Рекурсивно помечает все достижимые клетки региона как посещённые
func visit(Positiones map[Position]bool, p Position) {
	visited, exists := Positiones[p]
	if !exists {
		return // если клетки нет в регионе, выходим
	}
	if visited {
		return // если уже посещена, выходим
	}

	Positiones[p] = true // помечаем как посещённую

	nextPositiones := getNextMove(p)
	for _, nextPosition := range nextPositiones {
		visit(Positiones, nextPosition)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Process(in, out)
}
