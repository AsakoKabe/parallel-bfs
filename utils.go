package bfs

func createCubeGraph(size int) map[int][]int {
	graph := make(map[int][]int)

	// Функция для вычисления номера вершины в графе на основе координат
	vertex := func(x, y, z int) int {
		return x*size*size + y*size + z
	}

	// Добавляем рёбра для каждой вершины
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			for z := 0; z < size; z++ {
				v := vertex(x, y, z)
				var neighbors []int

				// Проверяем возможных соседей
				if x > 0 {
					neighbors = append(neighbors, vertex(x-1, y, z))
				}
				if x < size-1 {
					neighbors = append(neighbors, vertex(x+1, y, z))
				}
				if y > 0 {
					neighbors = append(neighbors, vertex(x, y-1, z))
				}
				if y < size-1 {
					neighbors = append(neighbors, vertex(x, y+1, z))
				}
				if z > 0 {
					neighbors = append(neighbors, vertex(x, y, z-1))
				}
				if z < size-1 {
					neighbors = append(neighbors, vertex(x, y, z+1))
				}

				graph[v] = neighbors
			}
		}
	}

	return graph
}
func checkBFSPath(bfsResult []int, graph map[int][]int, size int) bool {
	if len(bfsResult) < len(graph) {
		return false
	}

	visitedNodes := make(map[int]bool)
	for _, node := range bfsResult {
		visitedNodes[node] = true
	}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			for z := 0; z < size; z++ {
				node := x*size*size + y*size + z
				if !visitedNodes[node] {
					return false
				}
			}
		}
	}

	return true
}
