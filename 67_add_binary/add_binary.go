package add_binary

func addBinary(a string, b string) string {
	num := make([]byte, 0)

	carry := 0
	p1, p2 := len(a)-1, len(b)-1

	for p1 >= 0 || p2 >= 0 || carry > 0 {
		if p1 >= 0 {
			carry += int(a[p1] - '0')
			p1--
		}

		if p2 >= 0 {
			carry += int(b[p2] - '0')
			p2--
		}

		num = append(num, byte(carry%2+'0'))
		carry /= 2
	}

	for i := 0; i < len(num)/2; i++ {
		num[i], num[len(num)-i-1] = num[len(num)-i-1], num[i]
	}

	return string(num)
}
