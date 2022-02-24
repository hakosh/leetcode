package license_key_formatting

func licenseKeyFormatting(s string, k int) string {
	key := make([]byte, 0, len(s))
	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		char := s[i]

		if s[i] == '-' {
			continue
		} else if char >= 'a' && char <= 'z' {
			char = char - 'a' + 'A'
		}

		key = append(key, char)
		count++

		if count%k == 0 {
			key = append(key, '-')
		}
	}

	for len(key) > 0 && key[len(key)-1] == '-' {
		key = key[:len(key)-1]
	}

	for i := 0; i < len(key)/2; i++ {
		key[i], key[len(key)-1-i] = key[len(key)-1-i], key[i]
	}

	return string(key)
}
