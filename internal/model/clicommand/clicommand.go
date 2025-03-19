package clicommand

import "errors"

type CLICommand struct {
	Name        string
	Description string
	Execute     func(params ...string) error
}

var commands map[string]CLICommand = map[string]CLICommand{}

func Init() {
	commands = map[string]CLICommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Execute:     commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Execute:     commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Display the map or the next map",
			Execute:     commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Display the map or the previous map",
			Execute:     commandMapBack,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore the current map",
			Execute:     commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a Pokemon",
			Execute:     commandCatch,
		},
	}
}

func GetCommand(text string) (CLICommand, error) {
	command, ok := commands[text]
	if !ok {
		return CLICommand{}, errors.New("unknown command")
	}

	return command, nil

}
