package bfs

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func ParallelBFS(graph map[int][]int, source int, numThreads int) []int32 {
	visited := make([]int32, len(graph))

	var wg sync.WaitGroup
	currentFrontier := []int{source}
	visited[source] = 1

	for len(currentFrontier) > 0 {
		nextFrontier := make([]int, 0)
		var frontierMutex sync.Mutex
		chunkSize := (len(currentFrontier) + numThreads - 1) / numThreads
		for i := 0; i < numThreads; i++ {
			start := i * chunkSize
			end := min(start+chunkSize, len(currentFrontier))
			if start >= end {
				break
			}
			subList := currentFrontier[start:end]
			wg.Add(1)
			go func(batch []int) {
				defer wg.Done()
				localNextFrontier := make([]int, 0)
				for _, node := range batch {
					neighbors := graph[node]
					for _, neighbor := range neighbors {
						if atomic.CompareAndSwapInt32(&visited[neighbor], 0, 1) {
							localNextFrontier = append(localNextFrontier, neighbor)
						}
					}
				}
				frontierMutex.Lock()
				nextFrontier = append(nextFrontier, localNextFrontier...)
				frontierMutex.Unlock()
			}(subList)
		}

		wg.Wait()

		currentFrontier = nextFrontier
	}

	return visited
}

func TestParBFSPerformance(t *testing.T) {
	const NumTests = 5
	const size = 200
	const maxThreads = 4

	graph := createCubeGraph(size)
	startNode := 0

	var durations []time.Duration
	var totalDuration time.Duration

	for i := 0; i < NumTests; i++ {
		start := time.Now()
		visited := ParallelBFS(graph, startNode, maxThreads)
		duration := time.Since(start)

		var path []int
		for k, _ := range visited {
			path = append(path, k)
		}
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
