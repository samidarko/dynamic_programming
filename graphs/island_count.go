package graphs

import "fmt"

func IslandCount(grid [][]string) int {
	rowsCount := len(grid)
	colsCount := len(grid[0]) // assumed grid will never be of length 0
	visited := map[string]bool{}
	count := 0
	for rowIndex := 0; rowIndex < rowsCount; rowIndex++ {
		for colIndex := 0; colIndex < colsCount; colIndex++ {
			if IsExploreIsland(grid, rowIndex, colIndex, visited) {
				count++
			}
		}
	}
	return count
}

func IsExploreIsland(grid [][]string, rowIndex, colIndex int, visited map[string]bool) bool {
	rowInbounds := 0 <= rowIndex && rowIndex < len(grid)
	colInbounds := 0 <= colIndex && colIndex < len(grid[0])
	if !rowInbounds || !colInbounds {
		return false
	}

	if grid[rowIndex][colIndex] == "W" {
		return false
	}

	cell := fmt.Sprintf("%d,%d", rowIndex, colIndex)
	if visited[cell] {
		return false
	}

	visited[cell] = true

	IsExploreIsland(grid, rowIndex+1, colIndex, visited)
	IsExploreIsland(grid, rowIndex-1, colIndex, visited)
	IsExploreIsland(grid, rowIndex, colIndex+1, visited)
	IsExploreIsland(grid, rowIndex, colIndex-1, visited)

	return true
}
