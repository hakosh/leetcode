package reverse_bits

import (
	"math"
)

func reverseBits(num uint32) uint32 {
	var ans uint32

	for i := 31; i >= 0; i-- {
		if num%2 == 1 {
			ans += uint32(math.Pow(2, float64(i)))
		}

		num = num >> 1
	}

	return ans
}

func reverseBits2(num uint32) uint32 {
	ret := uint32(0)
	power := uint32(31)
	for num != 0 {
		ret += (num & 1) << power
		num = num >> 1
		power -= 1
	}
	return ret
}
