package main

import (
	"fmt"
)

// Brief solution description (in Russian):
//
// Короткое описание решения:
// * города -- вершины графа, дороги -- рёбра (неориентированные), склады могут размещаться в вершинах;
// * граф в памяти храним в виде списка смежных вершин (сет сетов);
// * каждый тест из входных данных прогоняем итеративно (не больше k итераций);
// * на каждой итерации находим вершину с наибольшим числом смежных вершин, размещаем в ней "склад", помечаем вершину
//   и все смежные ей как "отмеченные", удаляем вершину из графа;
// * если в какой момент не осталось непомеченных вершин в графе -- заканчиваем итеративный процесс досрочно
//   (используем не все возможные склады, а меньше).
//
// Сложность решения для каждого тест-кейса: O(k*n^2)

func main() {
	var testCount int
	fmt.Scanln(&testCount)
	ansStr := ""

	for t := 0; t < testCount; t++ {
		var n, m, k int
		fmt.Scanln(&n, &m, &k)

		// Build graphs.

		citiesCovered := make(map[int]bool, n)
		citiesNeighbours := make(map[int]map[int]bool) // a : b, bool "if active"

		for i := 1; i < n; i++ {
			citiesCovered[i] = false
		}

		for i := 0; i < m; i++ {
			var u, v int // 1-indexed
			fmt.Scanln(&u, &v)

			if len(citiesNeighbours[u]) == 0 {
				citiesNeighbours[u] = make(map[int]bool)
			}
			citiesNeighbours[u][v] = true

			if len(citiesNeighbours[v]) == 0 {
				citiesNeighbours[v] = make(map[int]bool)
			}
			citiesNeighbours[v][u] = true
		}

		ans := make([]int, 0, k)

		for k > 0 {
			mostLinkedNode := 0
			mostLinkedSize := -1
			mostLinkedAlreadyCovered := false

			// Find the most linked node (not covered and with those not covered).
			for node, neighbours := range citiesNeighbours {
				if citiesCovered[node] {
					continue
				}
				activeNeighbours := 0
				for neighbourNode, active := range neighbours {
					if active && !citiesCovered[neighbourNode] {
						activeNeighbours++
					}
				}
				if activeNeighbours > mostLinkedSize ||
					(activeNeighbours == mostLinkedSize && citiesCovered[node] == false) {
					mostLinkedNode = node
					mostLinkedSize = activeNeighbours
					mostLinkedAlreadyCovered = citiesCovered[node]
				}
			}

			if mostLinkedSize == -1 {
				break
			}
			if mostLinkedSize == 0 && mostLinkedAlreadyCovered {
				break
			}

			// Found the most linked => placing a warehouse there.

			ans = append(ans, mostLinkedNode)
			citiesCovered[mostLinkedNode] = true

			// Remove from graph.

			for neighbour := range citiesNeighbours[mostLinkedNode] {
				citiesNeighbours[neighbour][mostLinkedNode] = false
				citiesNeighbours[mostLinkedNode][neighbour] = false

				// All neighbours of covered node are also covered.
				citiesCovered[neighbour] = true
			}

			k--
		}

		ansStr += fmt.Sprintln(len(ans))
		for _, a := range ans {
			ansStr += fmt.Sprintf("%d ", a)
		}
		ansStr += fmt.Sprintln()
	}

	fmt.Print(ansStr)
}
