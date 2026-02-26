package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
}

func cleanInput(text string) []string {
	lowText := strings.ToLower(text)
	words := strings.Fields(lowText)
	return words
}

func StartRepl() {
	pagination := config{
		Next:     "",
		Previous: "",
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")

		scanner.Scan()

		cleanedInput := cleanInput(scanner.Text())[0]

		cmd, ok := getCommands()[cleanedInput]
		if !ok {
			fmt.Println("Uknonwn command")
			continue
		}
		if err := cmd.callback(&pagination); err != nil {
			fmt.Println(err)
		}
	}
}

func getCommands() map[string]cliCommand {
	commandsRegistry := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations in Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations when used after 'map' command",
			callback:    commandMapb,
		},
	}
	return commandsRegistry
}
