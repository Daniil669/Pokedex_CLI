package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "             Hello,        World!    ",
			expected: []string{"hello,", "world!"},
		},
		{
			input:    "one two three",
			expected: []string{"one", "two", "three"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    " word   ",
			expected: []string{"word"},
		},
	}
	for i, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Test #%d failed.\nexpected length: %d, actual length: %d\n", i+1, len(c.expected), len(actual))
			return
		}
		for j := range actual {
			word := actual[j]
			expectedWord := c.expected[j]
			if word != expectedWord {
				t.Errorf("expected word: %s, actual word: %s\n", expectedWord, word)
				return
			}
		}
	}
}
