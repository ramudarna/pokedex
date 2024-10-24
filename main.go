package main

import (
	"time"

	"github.com/ramudarna/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, time.Minute*5),
	}
	startRepl(&cfg)
}
