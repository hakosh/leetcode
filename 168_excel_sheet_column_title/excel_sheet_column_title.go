package excel_sheet_column_title

func convertToTitle(columnNumber int) string {
	if columnNumber == 0 {
		return "Z"
	} else if columnNumber <= 26 {
		return string(byte(columnNumber - 1 + 'A'))
	}

	return convertToTitle((columnNumber-1)/26) + convertToTitle(columnNumber%26)
}

func convertToTitleIterative(columnNumber int) string {
	reverse := make([]byte, 0)

	for columnNumber > 0 {
		columnNumber--
		reverse = append(reverse, byte(columnNumber%26+'A'))
		columnNumber /= 26
	}

	for i := 0; i < len(reverse)/2; i++ {
		reverse[i], reverse[len(reverse)-1-i] = reverse[len(reverse)-1-i], reverse[i]
	}

	return string(reverse)
}
