package main

import (
	"fmt"
	"pokedex_cli/repl"
)

func main() {
	words := repl.CleanInput("     Hello,     World!    ")
	for i, word := range words {
		fmt.Printf("%d: %s\n", i, word)
	}
}
