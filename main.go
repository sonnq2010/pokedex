package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sonnq2010/pokedexcli/internal/model/clicommand"
	stringutil "github.com/sonnq2010/pokedexcli/internal/pkg/util"
)

func main() {
	clicommand.Init()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := scanner.Text()

		words := stringutil.CleanInput(input)
		if len(words) == 0 {
			continue
		}

		firstWord := words[0]
		params := words[1:]

		command, err := clicommand.GetCommand(firstWord)
		if err != nil {
			fmt.Println("Unknown command")
			continue
		}

		err = command.Execute(params...)
		if err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}
