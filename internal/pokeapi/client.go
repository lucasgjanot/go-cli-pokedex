package pokeapi

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/lucasgjanot/go-cli-pokedex/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func GetData(url string) ([]byte, error) {
	var zero []byte
	res, err := http.Get(url)
	if res.StatusCode == 404 {
		return zero, errors.New("Resourse not found")
	}
	if err != nil {
		return zero, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return zero, err
	}

	return data, nil
}
