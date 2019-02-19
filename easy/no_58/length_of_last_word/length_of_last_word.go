package length_of_last_word

func lengthOfLastWord(s string) int {
	startj := len(s) - 1

	for startj >= 0 {
		if s[startj] != ' ' {
			break
		}
		startj --
	}
	starti := startj
	for starti >= 0 {
		if s[starti] == ' ' {
			break
		}
		starti--
	}
	return startj - starti
}
