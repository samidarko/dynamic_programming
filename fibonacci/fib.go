package fibonacci

type Memo = map[int]int

func Fib(n int, memo Memo) int {
	if fib, found := memo[n]; found {
		return fib
	}

	if n <= 2 {
		return 1
	}

	memo[n] = Fib(n-1, memo) + Fib(n-2, memo)
	return memo[n]
}

func FibTab(n int) int {
	tab := make([]int, n+3, n+3)
	tab[1] = 1

	for i := 0; i <= n; i++ {
		tab[i+1] += tab[i]
		tab[i+2] += tab[i]
	}

	return tab[n]
}
