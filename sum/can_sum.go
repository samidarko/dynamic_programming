package sum

func CanSum(targetSum int, numbers []int, memo map[int]bool) bool {
	if canSum, found := memo[targetSum]; found {
		return canSum
	}
	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for _, number := range numbers {
		remainder := targetSum - number
		if CanSum(remainder, numbers, memo) {
			memo[targetSum] = true
			return true
		}
	}

	memo[targetSum] = false
	return false
}

func CanSumTab(targetSum int, numbers []int) bool {
	table := make([]bool, targetSum+1)
	table[0] = true
	for i := 0; i <= targetSum; i++ {
		if table[i] {
			for _, n := range numbers {
				currentTargetSum := i + n
				if currentTargetSum < len(table) {
					table[currentTargetSum] = true
				}
			}
		}
	}
	return table[targetSum]
}
