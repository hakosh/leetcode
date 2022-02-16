package car_pooling

func carPooling(trips [][]int, capacity int) bool {
	var seats [1001]int

	for _, trip := range trips {
		seats[trip[1]] += trip[0]
		seats[trip[2]] -= trip[0]
	}

	seatsTaken := 0

	for _, num := range seats {
		seatsTaken += num

		if seatsTaken > capacity {
			return false
		}
	}

	return true
}
