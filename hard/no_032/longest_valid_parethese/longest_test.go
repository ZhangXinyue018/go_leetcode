package longest_valid_parethese

import "testing"

func TestLongest(t *testing.T) {
	testPairs := []struct {
		s   string
		num int
	}{
		{")(((((()())()()))()(()))(", 22},
	}
	for _, testPair := range testPairs {
		if temp := longestValidParentheses(testPair.s); temp != testPair.num {
			t.Errorf("%s targeting on %d but get %d\n", testPair.s, testPair.num, temp)
		}
	}
}
