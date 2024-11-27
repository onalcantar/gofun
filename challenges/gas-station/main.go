package main

func canCompleteCircuit(gas []int, cost []int) int {
	total, currentGas, startIndex := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		currentGas += gas[i] - cost[i]

		if currentGas < 0 {
			startIndex = i + 1
			currentGas = 0
		}
	}

	if total < 0 {
		return -1
	}

	return startIndex
}