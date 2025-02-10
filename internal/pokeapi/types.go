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

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
