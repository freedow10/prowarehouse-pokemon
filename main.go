package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model "github.com/freedow10/prowarehouse-pokemon/Model"
	"github.com/freedow10/prowarehouse-pokemon/database"
	"github.com/gorilla/mux"
)

type pokemonViewModel struct {
	Name   string   `json:"name"`
	Weight float32  `json:"weight"`
	Height float32  `json:"height"`
	Moves  []string `json:"moves"`
	Types  []string `json:"types"`
}

func GetAListOfPokemon(w http.ResponseWriter, r *http.Request) {
	db := database.InitDatabase("./db/pokemon.db")
	w.Header().Set("Content-Type", "application/json")

	l, err := db.GetAListOfPokemonFromDB(20)
	if err != nil {
		fmt.Println(err)
	}

	pokemonList := []pokemonViewModel{}

	for _, elements := range l {
		pokemonList = append(pokemonList, convertPokemonModelToViewModel(elements))
	}

	dat, _ := json.Marshal(pokemonList)
	w.Write(dat)
}

func convertPokemonModelToViewModel(input model.Pokemon) pokemonViewModel {

	var tmpMoveSlice []string
	var tmpTypeSlice []string

	for i, m := range input.Moves {

		if i >= 4 {
			break
		}

		tmpMoveSlice = append(tmpMoveSlice, m)
	}

	for _, t := range input.Types {
		tmpTypeSlice = append(tmpTypeSlice, t)
	}

	return pokemonViewModel{
		Name:   input.Name,
		Weight: float32(input.Weight),
		Height: float32(input.Height),
		Moves:  tmpMoveSlice,
		Types:  tmpTypeSlice,
	}
}

func main() {

	// fmt.Println(db.EmptyTableData())

	// fmt.Println(db.FillDatabase(100))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", GetAListOfPokemon).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}
