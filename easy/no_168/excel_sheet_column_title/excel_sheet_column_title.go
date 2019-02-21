package excel_sheet_column_title

func convertToTitle(n int) string {
	result := []byte{}
	for n != 0 {
		temp := (n-1)%26 + 1
		result = append(result, 'A' + byte(temp) - 1)
		if n > 26 {
			n = (n - temp) / 26
		} else {
			break
		}
	}

	i := 0
	j := len(result) - 1
	for i < j {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}

	return string(result)
}
