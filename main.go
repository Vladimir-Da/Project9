package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	rnd := rand.NewSource(time.Now().Unix())
	r := rand.New(rnd)
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = r.Int()
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) <= CHUNKS {
		max := data[0]
		for _, v := range data[1:] {
			if v > max {
				max = v
			}
		}
		return max
	}
	leinghtChunk := len(data) / CHUNKS
	if leinghtChunk == 0 {
		leinghtChunk = 1
	}
	result := make([]int, 0, CHUNKS)
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		start := i * leinghtChunk
		if start >= len(data) {
			continue
		}
		end := start + leinghtChunk
		if i == CHUNKS-1 || end > len(data) {
			end = len(data)
		}
		go func(beg, fin int) {
			defer wg.Done()
			if beg >= fin {
				return
			}
			max := data[beg]
			for _, v := range data[beg+1 : fin] {
				if v > max {
					max = v
				}
			}
			mu.Lock()
			result = append(result, max)
			mu.Unlock()
		}(start, end)
	}
	wg.Wait()

	if len(result) == 0 {
		return 0
	}

	maxVal := result[0]
	for _, v := range result[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	now := time.Now()
	max := maximum(data)
	elapsed := time.Since(now).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	now = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(now).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
