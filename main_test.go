package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateRandomELenemts(t *testing.T) {
	data := generateRandomElements(10)
	require.Equal(t, 10, len(data))

	for _, value := range data {
		require.Greater(t, value, 0)
	}
	empty := generateRandomElements(0)
	require.Empty(t, empty)
}

func TestMaximum(t *testing.T) {
	require.Equal(t, 10, maximum([]int{3, 8, 2, 10, 5}))
	require.Equal(t, 7, maximum([]int{7}))
	require.Equal(t, 0, maximum([]int{}))
	require.Equal(t, -3, maximum([]int{-5, -10, -3}))
}

func TestMaxChunks(t *testing.T) {
	require.Equal(t, 10, maxChunks([]int{3, 8, 2, 10, 5, 7, 1, 9}))
	require.Equal(t, 7, maxChunks([]int{7}))
	require.Equal(t, 0, maximum([]int{}))
	require.Equal(t, -3, maxChunks([]int{-5, -10, -3}))
}
