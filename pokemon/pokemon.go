package pokemon

import (
	"strings"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetAPokemon(input string) (*structs.Pokemon, error) {
	l, err := pokeapi.Pokemon(input)
	if err != nil {
		return nil, err
	}

	return &l, nil
}

func FetchAListPokemon(lenghth int) ([]*structs.Pokemon, error) {
	var pokemonList = []*structs.Pokemon{}

	r, err := pokeapi.Resource("pokemon", 0, lenghth)
	if err != nil {
		return nil, err
	}

	for _, e := range r.Results {
		s := strings.Split(e.URL, "/")

		p, err := GetAPokemon(s[len(s)-2])
		if err != nil {
			return nil, err
		}

		pokemonList = append(pokemonList, p)
	}

	return pokemonList, nil
}
