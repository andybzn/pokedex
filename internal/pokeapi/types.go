package pokeapi

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationData struct {
	Count     int        `json:"count"`
	Next      *string    `json:"next"`
	Previous  *string    `json:"previous"`
	Locations []Location `json:"results"`
}

type LocationArea struct {
	Name       string             `json:"name"`
	Location   Location           `json:"location"`
	Encounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokedex struct {
	Pokemon map[string]Pokemon
}

type Pokemon struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		Name string `json:"name"`
	} `json:"abilities"`
	Forms []struct {
		Name string `json:"name"`
	} `json:"forms"`
	Moves []struct {
		Name string `json:"name"`
	} `json:"moves"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `jsoin:"name"`
		} `json:"stat"`
	} `json:"stats"`
}
