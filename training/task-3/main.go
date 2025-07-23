package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Relief [][]rune

func process(in *bufio.Reader, out *bufio.Writer) {
	var t int
	fmt.Fscanln(in, &t)
	
	for test := 0; test < t; test++ {
		var k, n, m int
		fmt.Fscanln(in, &k, &n, &m)
		reliefs := make([]Relief, k)
		for i := 0; i < k; i++ {
			reliefs[i] = make(Relief, n)
			for j := 0; j < n; j++ {
				line, _ := in.ReadString('\n')
				line = strings.TrimRight(line, "\r\n")
				reliefs[i][j] = []rune(line)
			}
			// Пропускаем пустую строку между рельефами, кроме последнего
			if i != k-1 {
				in.ReadString('\n')
			}
		}
		// Итоговый рельеф
		result := make([][]rune, n)
		for i := 0; i < n; i++ {
			result[i] = make([]rune, m)
			for j := 0; j < m; j++ {
				result[i][j] = '.'
			}
		}
		// Наложение рельефов
		for l := 0; l < k; l++ {
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					if reliefs[l][i][j] != '.' && result[i][j] == '.' {
						result[i][j] = reliefs[l][i][j]
					}
				}
			}
		}
		// Вывод результата
		for i := 0; i < n; i++ {
			out.WriteString(string(result[i]) + "\n")
		}
		out.WriteString("\n")
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	process(in, out)
}
