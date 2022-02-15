package median_of_two_sorted_arrays

import (
	"math"
)

// Find Kth Idea:
// https://leetcode.com/problems/median-of-two-sorted-arrays/discuss/295080/Go

// https://www.youtube.com/watch?v=CXoF9SkUq7E

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}

	short, long := nums1, nums2
	if len(short) > len(long) {
		short, long = long, short
	}

	direction := 1
	lShort, rShort, lLong, rLong := 1, 1, 1, 1

	left, right := 0, len(short)

	for direction != 0 {
		mid := (left + right) / 2

		lShort, rShort, lLong, rLong = getIndices(mid, short, long)
		direction = getDirection(lShort, rShort, lLong, rLong, short, long)

		if direction < 0 {
			right = mid - 1
		} else if direction > 0 {
			left = mid + 1
		}
	}

	return getResult(lShort, rShort, lLong, rLong, short, long)
}

func getIndices(rShort int, short, long []int) (int, int, int, int) {
	midIndex := (len(short) + len(long)) / 2
	rLong := midIndex - rShort

	return rShort - 1, rShort, rLong - 1, rLong
}

func getVal(arr []int, i int) int {
	if i == -1 {
		return math.MinInt32
	}

	if i == len(arr) {
		return math.MaxInt32
	}

	return arr[i]
}

func getDirection(lShort, rShort, lLong, rLong int, short, long []int) int {
	if getVal(short, lShort) > getVal(long, rLong) {
		return -1
	} else if getVal(long, lLong) > getVal(short, rShort) {
		return 1
	} else {
		return 0
	}
}

func getResult(lShort, rShort, lLong, rLong int, short, long []int) float64 {
	if (len(short)+len(long))%2 == 0 {
		a := float64(max(getVal(short, lShort), getVal(long, lLong)))
		b := float64(min(getVal(short, rShort), getVal(long, rLong)))

		return (a + b) / 2
	} else {
		return float64(min(getVal(long, rLong), getVal(short, rShort)))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
