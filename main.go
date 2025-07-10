package main

import (
	"fmt"
	"math/rand"
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
	leinghtChunk := len(data) / CHUNKS
	// if len(data)< CHUNKS
	if leinghtChunk == 0 {
		return maximum(data)
	}
	result := make(chan int)
	for i := 1; i <= CHUNKS; i++ {
		var chunk []int
		if i != CHUNKS {
			chunk = data[:leinghtChunk]
			data = data[leinghtChunk:]
			// last chunk(8)
		} else {
			chunk = data[:]
		}
		go func(chunk []int) {
			//writing to chanel result
			result <- maximum(chunk)
		}(chunk)
	}
	fin := make([]int, 0, CHUNKS)
	for range CHUNKS {
		//reading from chanel result
		fin = append(fin, <-result)
	}
	close(result)

	return maximum(fin)
}
func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	now := time.Now()
	max := maximum(data)
	elapsed := time.Since(now).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	now = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(now).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
