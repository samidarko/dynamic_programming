package graphs

import (
	"fmt"
	"math"
)

func MinimumIsland(grid [][]string) int {
	rowsCount := len(grid)
	colsCount := len(grid[0]) // assumed grid will never be of length 0
	visited := map[string]bool{}
	minimum := math.MaxInt
	for rowIndex := 0; rowIndex < rowsCount; rowIndex++ {
		for colIndex := 0; colIndex < colsCount; colIndex++ {
			islandSize := IslandSize(grid, rowIndex, colIndex, visited)
			if islandSize > 0 && islandSize < minimum {
				minimum = islandSize
			}
		}
	}
	return minimum
}

func IslandSize(grid [][]string, rowIndex, colIndex int, visited map[string]bool) int {
	rowInbounds := 0 <= rowIndex && rowIndex < len(grid)
	colInbounds := 0 <= colIndex && colIndex < len(grid[0])
	if !rowInbounds || !colInbounds {
		return 0
	}

	if grid[rowIndex][colIndex] == "W" {
		return 0
	}

	cell := fmt.Sprintf("%d,%d", rowIndex, colIndex)
	if visited[cell] {
		return 0
	}

	visited[cell] = true

	size := 1
	size += IslandSize(grid, rowIndex+1, colIndex, visited)
	size += IslandSize(grid, rowIndex-1, colIndex, visited)
	size += IslandSize(grid, rowIndex, colIndex+1, visited)
	size += IslandSize(grid, rowIndex, colIndex-1, visited)

	return size
}
