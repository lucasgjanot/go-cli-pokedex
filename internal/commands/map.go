package commands

import (
	"errors"
	"fmt"

	"github.com/lucasgjanot/go-cli-pokedex/internal/config"
)

type LocationArea struct{}

func CommandMapf(cfg *config.Config, args ...string) error {
	if len(args) != 0 {
		return errors.New("usage: map")
	}
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

func CommandMapb(cfg *config.Config, args ...string) error {
	if len(args) != 0 {
		fmt.Println("usage: mapb")
	}

	if cfg.PrevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.PokeapiClient.ListLocations(cfg.PrevLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationResp.Next
	cfg.PrevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
