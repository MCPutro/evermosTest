package Task2

import (
	"fmt"
	"testing"
)

var (
	territory = [][2]int{
		{-1, 0},           // atas
		{1, 0},            // bawah
		{0, -1},           // kiri
		{0, 1},            // Right
		{-1, -1}, {-1, 1}, // diagonal
		{1, -1}, {1, 1}, // Diagonals
	}
)

func checkTerritory(i, j, m, n int, grid [][]byte) {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '.' {
		return
	}

	// Tandai sel sebagai sudah dikunjungi
	grid[i][j] = '.'

	for _, dir := range territory {
		checkTerritory(i+dir[0], j+dir[1], m, n, grid)
	}
}

func count(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	c := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '#' {
				c++
				checkTerritory(i, j, m, n, grid)
			}
		}
	}

	return c
}

func TestTask2(t *testing.T) {
	grid := [][]byte{
		{'#', '.', '#', '#', '#'},
		{'.', '.', '#', '#', '#'},
		{'#', '.', '#', '#', '#'},
		{'#', '#', '#', '#', '.'},
	}

	fmt.Println(count(grid))
}
