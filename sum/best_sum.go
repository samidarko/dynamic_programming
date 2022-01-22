package sum

func BestSum(targetSum int, numbers []int, memo map[int][]int) []int {
	if howSum, found := memo[targetSum]; found {
		return howSum
	}

	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	var shortestResult []int = nil

	for _, number := range numbers {
		remainder := targetSum - number
		result := BestSum(remainder, numbers, memo)
		if result != nil {
			result = append(result, number)
			if shortestResult == nil || len(result) < len(shortestResult) {
				shortestResult = result
			}
		}
	}

	memo[targetSum] = shortestResult
	return shortestResult
}

func BestSumTab(targetSum int, numbers []int) []int {
	table := make([][]int, targetSum+1)
	table[0] = make([]int, 0)

	for i := 0; i <= targetSum; i++ {
		if table[i] != nil {
			for _, n := range numbers {
				currentTargetSum := i + n
				howSum := append(table[i], n)
				if currentTargetSum < len(table) && (table[currentTargetSum] == nil || len(howSum) < len(table[currentTargetSum])) {
					table[currentTargetSum] = howSum
				}
			}
		}
	}

	return table[targetSum]
}
