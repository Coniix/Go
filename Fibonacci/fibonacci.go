package fibonacci

func FibRecursion(n int) []int {
	fibSequence := []int{}

	if n == 0 {
		fibSequence = append(fibSequence, 0)
		return fibSequence
	} else {
		fibSequence = FibRecursion(n - 1)
	}

	if n == 1 {
		fibSequence = append(fibSequence, 1)
	} else {
		fibSequence = append(fibSequence, (fibSequence[n-2] + fibSequence[n-1]))
	}

	return fibSequence
}

func FibLoop(n int) []int {
	fibSequence := []int{}

	for i := 0; i <= n; i++ {
		if i == 0 {
			fibSequence = append(fibSequence, i)
			continue
		} else if i == 1 {
			fibSequence = append(fibSequence, i)
			continue
		}

		fibSequence = append(fibSequence, fibSequence[i-2]+fibSequence[i-1])
	}

	return fibSequence
}
