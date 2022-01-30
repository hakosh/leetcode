package interleaving_string

var cache [][]int

func isInterleave(s1 string, s2 string, s3 string) bool {
	cache = make([][]int, len(s1)+1)
	for i := range cache {
		cache[i] = make([]int, len(s2)+1)
	}

	if len(s1)+len(s2) != len(s3) {
		return false
	}

	ans := dp(s1, s2, s3, 0, 0)

	if ans == 1 {
		return true
	} else {
		return false
	}
}

func dp(s1, s2, s3 string, ptr1, ptr2 int) int {
	ptr3 := ptr1 + ptr2

	if ptr3 == len(s3) {
		return 1
	}

	if cache[ptr1][ptr2] == 0 {
		ans := 2

		if len(s1) == ptr1 {
			if s2[ptr2] == s3[ptr3] {
				ans = dp(s1, s2, s3, ptr1, ptr2+1)
			} else {
				ans = 2
			}
		} else if len(s2) == ptr2 {
			if s1[ptr1] == s3[ptr3] {
				ans = dp(s1, s2, s3, ptr1+1, ptr2)
			} else {
				ans = 2
			}
		} else {
			if s1[ptr1] == s2[ptr2] && s1[ptr1] == s3[ptr3] {
				a := dp(s1, s2, s3, ptr1+1, ptr2) == 1
				b := dp(s1, s2, s3, ptr1, ptr2+1) == 1

				if a || b {
					ans = 1
				} else {
					ans = 2
				}
			} else if s1[ptr1] == s3[ptr3] {
				ans = dp(s1, s2, s3, ptr1+1, ptr2)
			} else if s2[ptr2] == s3[ptr3] {
				ans = dp(s1, s2, s3, ptr1, ptr2+1)
			}
		}

		cache[ptr1][ptr2] = ans
	}

	return cache[ptr1][ptr2]
}
