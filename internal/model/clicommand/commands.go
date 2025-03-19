package clicommand

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/sonnq2010/pokedexcli/internal/model"
	pokerepository "github.com/sonnq2010/pokedexcli/internal/repository/poke_repository"
)

var currentLocation = model.LocationResponse{}
var caughtPokemons = map[string]model.Pokemon{}

func commandHelp(params ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}

func commandExit(params ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(params ...string) error {
	repo := pokerepository.PokeRepository{}

	next := currentLocation.Next

	location, err := repo.GetLocation(next)
	if err != nil {
		return err
	}

	currentLocation = location

	names := location.GetLocationNames()
	for _, name := range names {
		fmt.Println(name)
	}

	return nil
}

func commandMapBack(params ...string) error {
	repo := pokerepository.PokeRepository{}

	prev := currentLocation.Previous
	if prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	location, err := repo.GetLocation(prev)
	if err != nil {
		return err
	}

	currentLocation = location

	names := location.GetLocationNames()
	for _, name := range names {
		fmt.Println(name)
	}

	return nil
}

func commandExplore(params ...string) error {
	if len(params) == 0 {
		fmt.Println("Please specify a location to explore")
		return nil
	}

	area := params[0]
	fmt.Printf("Exploring %s...\n", area)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", area)
	repo := pokerepository.PokeRepository{}

	locationDetail, err := repo.GetLocationDetail(url)
	if err != nil {
		return err
	}

	pokemons := locationDetail.GetPokemonNames()
	if len(pokemons) == 0 {
		fmt.Println("No pokemons found in this location")
		return nil
	}

	fmt.Println("Pokemons found:")
	for _, p := range pokemons {
		fmt.Println(" -", p)
	}
	return nil
}

func commandCatch(params ...string) error {
	if len(params) == 0 {
		fmt.Println("Please specify a Pokemon to catch")
		return nil
	}

	pokemonName := params[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemonName)

	repo := pokerepository.PokeRepository{}
	pokemon, err := repo.GetPokemonDetail(url)
	if err != nil {
		return err
	}

	if pokemon.Name == "" {
		fmt.Println("Pokemon not found")
		return nil
	}

	baseExp := float64(pokemon.BaseExperience) / 255
	threshold := 70 - int(50*baseExp) // Adjust these numbers based on desired difficulty
	if rand.Intn(100) < threshold {
		fmt.Printf("%s was caught!\n", pokemonName)
		caughtPokemons[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
