package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Process(input *bufio.Reader, output *bufio.Writer) {
	var t int
	fmt.Fscan(input, &t)
	input.ReadString('\n')

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(input, &n)
		input.ReadString('\n') // Считываем оставшийся символ новой строки

		scores := make(map[string]int)      // Словарь для хранения очков каждого участника
		var names []string                  // Список всех уникальных имён (для сортировки)
		nameExists := make(map[string]bool) // Для проверки уникальности имён
		var action string                   // Для хранения действия x

		// Обрабатываем n строк событий
		for i := 0; i < n; i++ {
			line, _ := input.ReadString('\n')
			line = strings.TrimSpace(line)
			parts := strings.Fields(line)

			if len(parts) == 0 {
				continue
			}

			// Получаем имя говорящего (убираем '!')
			a := parts[0][:len(parts[0])-1]
			if !nameExists[a] {
				nameExists[a] = true
				names = append(names, a)
			}

			// Определяем действие x (оно всегда на одном и том же месте)
			if action == "" {
				if parts[1] == "I" && len(parts) >= 4 && parts[2] == "am" {
					if len(parts) == 5 && parts[3] == "not" {
						action = parts[4][:len(parts[4])-1]
					} else {
						action = parts[3][:len(parts[3])-1]
					}
				} else if len(parts) >= 4 {
					if len(parts) == 5 && parts[3] == "not" {
						action = parts[4][:len(parts[4])-1]
					} else {
						action = parts[3][:len(parts[3])-1]
					}
				}
			}

			if parts[1] == "I" && len(parts) >= 4 && parts[2] == "am" {
				// Формат: X! I am ...
				if len(parts) == 5 && parts[3] == "not" {
					scores[a]-- // "Я не ..." — минус балл
				} else {
					scores[a] += 2 // "Я ..." — плюс два балла
				}
			} else if len(parts) >= 4 {
				// Формат: X! Y is ...
				b := parts[1]
				if !nameExists[b] {
					nameExists[b] = true
					names = append(names, b)
				}
				if len(parts) == 5 && parts[3] == "not" {
					scores[b]-- // "Y не ..." — минус балл
				} else {
					scores[b]++ // "Y ..." — плюс балл
				}
			}
		}

		// Находим максимальный балл
		maxScore := -1 << 63
		for _, name := range names {
			if scores[name] > maxScore {
				maxScore = scores[name]
			}
		}

		// Собираем всех кандидатов с максимальным баллом
		var candidates []string
		for _, name := range names {
			if scores[name] == maxScore {
				candidates = append(candidates, name)
			}
		}
		sort.Strings(candidates) // Сортируем по алфавиту

		// Выводим имена победителей в нужном формате
		for _, name := range candidates {
			fmt.Fprintf(output, "%s is %s.\n", name, action)
		}
	}
}

func main() {
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	Process(input, output)
}
