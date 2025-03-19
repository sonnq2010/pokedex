package pokerepository

import (
	"encoding/json"
	"net/http"

	"github.com/sonnq2010/pokedexcli/internal/model"
)

type PokeRepository struct{}

func (r *PokeRepository) GetLocation(next string) (model.LocationResponse, error) {
	url := next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location"
	}

	resp, err := http.Get(url)
	if err != nil {
		return model.LocationResponse{}, err
	}
	defer resp.Body.Close()

	locationResp := model.LocationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&locationResp); err != nil {
		return model.LocationResponse{}, err
	}

	return locationResp, nil
}

func (r *PokeRepository) GetLocationDetail(url string) (model.LocationDetail, error) {
	resp, err := http.Get(url)
	if err != nil {
		return model.LocationDetail{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.LocationDetail{}, nil
	}

	locationDetail := model.LocationDetail{}
	if err := json.NewDecoder(resp.Body).Decode(&locationDetail); err != nil {
		return model.LocationDetail{}, err
	}

	return locationDetail, nil
}

func (r *PokeRepository) GetPokemonDetail(url string) (model.Pokemon, error) {
	resp, err := http.Get(url)
	if err != nil {
		return model.Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Pokemon{}, nil
	}

	pokemon := model.Pokemon{}
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return model.Pokemon{}, err
	}

	return pokemon, nil
}
