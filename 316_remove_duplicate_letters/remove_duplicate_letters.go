package remove_duplicate_letters

func removeDuplicateLetters(s string) string {
	last := make([]int, 127)
	for i, _ := range s {
		last[s[i]] = i
	}

	has := make([]bool, 127)
	ans := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if !has[s[i]] {
			for len(ans) > 0 && s[i] < ans[len(ans)-1] && last[ans[len(ans)-1]] > i {
				top := ans[len(ans)-1]
				has[top] = false
				ans = ans[:len(ans)-1]
			}

			has[s[i]] = true
			ans = append(ans, s[i])
		}
	}

	return string(ans)
}
