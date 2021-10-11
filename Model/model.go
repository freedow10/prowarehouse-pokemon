package model

import "strings"

type Pokemon struct {
	Name   string
	Weight float32
	Height float32
	Moves  []string
	Types  []string
}

func GetTestPokemon() Pokemon {

	testPokemon := Pokemon{
		Name:   "testPokemon",
		Weight: 1.000,
		Height: 1.000,
		Moves:  []string{"Move1", "Move2", "Move3"},
		Types:  []string{"Type1", "Type2", "Type3"},
	}

	return testPokemon
}

func (p Pokemon) GetPokemonMoves() string {
	return strings.Join(p.Moves, ",")
}

func (p Pokemon) GetPokemonTypes() string {
	return strings.Join(p.Types, ",")

}
