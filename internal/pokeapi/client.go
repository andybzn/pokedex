package pokeapi

import (
	"net/http"
	"time"

	"github.com/andybzn/pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, cacheTimeout time.Duration) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
		cache:      pokecache.NewCache(cacheTimeout),
	}
}
