package repl

import (
	"strings"
)

func CleanInput(text string) []string {
	lowText := strings.ToLower(text)
	words := strings.Fields(lowText)
	return words
}
