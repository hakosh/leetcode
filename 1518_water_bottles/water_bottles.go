package water_bottles

func numWaterBottles(numBottles int, numExchange int) int {
	drank := 0
	full, empty := numBottles, 0

	for full != 0 || empty > numExchange-1 {
		drank += full
		empty += full

		full = empty / numExchange
		empty = empty % numExchange
	}

	return drank
}
