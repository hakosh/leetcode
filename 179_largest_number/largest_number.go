package largest_number

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	numsStr := make([]string, len(nums))
	for i, num := range nums {
		numsStr[i] = strconv.Itoa(num)
	}

	sort.Slice(numsStr, func(i, j int) bool {
		order1 := numsStr[i] + numsStr[j]
		order2 := numsStr[j] + numsStr[i]

		return strings.Compare(order1, order2) == 1
	})

	number := strings.Join(numsStr, "")
	if number[0] == '0' {
		return "0"
	} else {
		return number
	}
}
