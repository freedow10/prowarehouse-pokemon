package app

import (
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/mock"
)

type MockPokemoner struct {
	mock.Mock
}

func NewMockPokemoner() *MockPokemoner {
	return &MockPokemoner{}
}

func (mpk *MockPokemoner) FetchAListPokemon(lenghth int) ([]*structs.Pokemon, error) {
	args := mpk.Called(lenghth)

	return args[0].([]*structs.Pokemon), args.Error(1)
}

func (mpk *MockPokemoner) GetAPokemon(input string) (*structs.Pokemon, error) {
	args := mpk.Called(input)

	return args[0].(*structs.Pokemon), args.Error(1)
}
