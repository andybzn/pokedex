package pokeapi

type Config struct {
	BaseURL    string
	APIVersion string
}

var APIConfig = Config{
	BaseURL:    "https://pokeapi.co/api/",
	APIVersion: "v2",
}
