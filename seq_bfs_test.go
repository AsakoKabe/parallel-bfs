package bfs

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
)

func seqBFS(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	queue := []int{start}
	var result []int

	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

func TestSeqBFSPerformance(t *testing.T) {
	const NumTests = 5
	const size = 200

	graph := createCubeGraph(size)
	startNode := 0

	var durations []time.Duration
	var totalDuration time.Duration

	for i := 0; i < NumTests; i++ {
		start := time.Now()
		path := seqBFS(graph, startNode)
		duration := time.Since(start)

		assert.True(t, checkBFSPath(path, graph, size))

		durations = append(durations, duration)
		totalDuration += duration
	}

	averageDuration := totalDuration / 5

	fmt.Println("Run\tExecution Time")
	fmt.Println("------------------------")
	for i, duration := range durations {
		fmt.Printf("%d\t%v\n", i+1, duration)
	}
	fmt.Println("------------------------")
	fmt.Printf("Average Execution Time: %v\n", averageDuration)
}
