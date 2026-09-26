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
	// ваш код здесь
	if size <= 0 {
		return []int{}
	}
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(size) + 1
	}

	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, value := range data {
		if value > max {
			max = value
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	var wg sync.WaitGroup
	chunkSize := (len(data) + CHUNKS - 1) / CHUNKS
	maxValues := make([]int, min(len(data), CHUNKS))
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}
		if start >= len(data) {
			break
		}
		wg.Add(1)

		go func(i, start, end int) {
			defer wg.Done()
			maxValues[i] = maximum(data[start:end])
		}(i, start, end)
	}
	wg.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	startElement := time.Now()
	maxElement := maximum(data)
	elapsedElement := time.Since(startElement)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxElement, elapsedElement.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	startChunks := time.Now()
	maxChunk := maxChunks(data)
	elapsedChunks := time.Since(startChunks)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxChunk, elapsedChunks.Microseconds())
}
