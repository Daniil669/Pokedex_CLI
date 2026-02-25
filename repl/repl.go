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
	callback    func() error
}

type config struct {
}

func cleanInput(text string) []string {
	lowText := strings.ToLower(text)
	words := strings.Fields(lowText)
	return words
}

func StartRepl() {
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
		if err := cmd.callback(); err != nil {
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
			description: "Displays 20 locations in Pokemon world. If called for the 2nd and so on times, it will display next 20 locations",
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
