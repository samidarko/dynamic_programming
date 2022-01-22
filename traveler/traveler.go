package traveler

import "fmt"

func getKey(n, m int) string {
	if n < m {
		return fmt.Sprintf("%d:%d", n, m)
	}
	return fmt.Sprintf("%d:%d", m, n)
}

func GridTraveler(m, n int, memo map[string]int) int {
	key := getKey(n, m)
	if value, found := memo[key]; found {
		return value
	}
	if n == 0 || m == 0 {
		return 0
	}
	if n == 1 || m == 1 {
		return 1
	}
	memo[key] = GridTraveler(m-1, n, memo) + GridTraveler(m, n-1, memo)
	return memo[key]
}

func GridTravelerTab(m, n int) int {
	table := make([][]int, m+1)
	for i := range table {
		table[i] = make([]int, n+1)
	}
	table[1][1] = 1
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			current := table[i][j]
			if j+1 <= n {
				table[i][j+1] += current
			}
			if i+1 <= m {
				table[i+1][j] += current
			}
		}
	}

	return table[m][n]
}
