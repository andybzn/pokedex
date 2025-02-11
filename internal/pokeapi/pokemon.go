package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchPokemon(name string) (Pokemon, error) {
	pokemonUrl := APIConfig.BaseURL + APIConfig.APIVersion + "/pokemon/" + name

	data, exists := c.cache.Get(pokemonUrl)
	if !exists {
		req, err := http.NewRequest("GET", pokemonUrl, nil)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error forming request: %v", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error fetching data: %v", err)
		}
		defer res.Body.Close()

		resData, err := io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error reading response: %v", err)
		}

		c.cache.Add(pokemonUrl, resData)
		data = resData
	}

	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, fmt.Errorf("failed to parse pokemon: %v", err)
	}

	return pokemon, nil
}
