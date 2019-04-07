package restore_ip_addresses

func restoreIpAddresses(s string) []string {
	result := &[]string{}
	DetailRestore("", s, 4, result)
	return *result
}

func DetailRestore(preStr, s string, n int, result *[]string) {
	if n == 0 && len(s) == 0 {
		*result = append(*result, preStr[0:len(preStr)-1])
		return
	} else if len(s) == 0 || n <= 0 {
		return
	} else if s[0] == '0' {
		DetailRestore(preStr+"0.", s[1:], n-1, result)
		return
	}
	i := 0
	var temp string
	for i < len(s) {
		i++
		temp = string(s[:i])
		if i > 3 || (i == 3 && temp > "255") {
			break
		}
		DetailRestore(preStr+string(temp)+".", s[i:], n-1, result)
	}
}
