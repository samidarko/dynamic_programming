package sum

func HowSum(targetSum int, numbers []int, memo map[int][]int) []int {
	if howSum, found := memo[targetSum]; found {
		return howSum
	}
	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	for _, number := range numbers {
		remainder := targetSum - number
		result := HowSum(remainder, numbers, memo)
		if result != nil {
			memo[targetSum] = append(result, number)
			return memo[targetSum]
		}
	}

	memo[targetSum] = nil
	return nil
}

func HowSumTab(targetSum int, numbers []int) []int {
	table := make([][]int, targetSum+1)
	table[0] = make([]int, 0)
	for i := 0; i <= targetSum; i++ {
		if table[i] != nil {
			for _, n := range numbers {
				currentTargetSum := i + n
				if currentTargetSum < len(table) {
					table[currentTargetSum] = append(table[i], n)
				}
			}
		}
	}
	return table[targetSum]
}
