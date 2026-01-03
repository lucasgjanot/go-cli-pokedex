package commands

import (
	"errors"
	"fmt"

	"github.com/lucasgjanot/go-cli-pokedex/internal/config"
)

type LocationArea struct{}

func CommandMapf(cfg *config.Config) error {
	if cfg.PrevLocationsURL != nil && cfg.NextLocationsURL == nil {
		return errors.New("you're on the last page")
	}
	locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.NextLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationsResp.Next
	cfg.PrevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil

}

func CommandMapb(cfg *config.Config) error {
	if cfg.PrevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.PrevLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationsResp.Next
	cfg.PrevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
