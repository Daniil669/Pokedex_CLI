package repl

import "fmt"

func commandHelp(c *config) error {
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	commands := getCommands()
	for _, s := range commands {
		fmt.Printf("%s: %s\n", s.name, s.description)
	}
	fmt.Printf("\n")
	return nil
}
