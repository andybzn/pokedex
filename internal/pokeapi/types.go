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
