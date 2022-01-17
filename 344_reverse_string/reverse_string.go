package reverse_string

func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}

	end := len(s) - 1
	s[0], s[end] = s[end], s[0]

	reverseString(s[1:end])
}
