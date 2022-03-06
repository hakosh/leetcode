package champagne_tower

func champagneTower(poured int, query_row int, query_glass int) float64 {
	this := []float64{float64(poured)}
	next := []float64{0, 0}

	for i := 0; i <= query_row; i++ {
		for j := range this {
			if this[j] > 1 {
				overflow := this[j] - 1
				this[j] = 1

				next[j] += overflow / 2
				next[j+1] += overflow / 2
			}
		}

		if i < query_row {
			this, next = next, make([]float64, len(next)+1)
		}
	}

	return this[query_glass]
}
