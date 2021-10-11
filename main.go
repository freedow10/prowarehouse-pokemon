package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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

	l, err := db.GetAListOfPokemonFromDB(500)
	if err != nil {
		fmt.Println(err)
	}

	pokemonList := []pokemonViewModel{}

	for _, elements := range l {
		pokemonList = append(pokemonList, convertPokemonModelToViewModel(elements))
	}

	jsonPokemon, _ := json.Marshal(pokemonList)
	w.Write(jsonPokemon)
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
		Name:   strings.Title(input.Name),
		Weight: input.Weight,
		Height: input.Height,
		Moves:  tmpMoveSlice,
		Types:  tmpTypeSlice,
	}
}

func main() {
	// db := database.InitDatabase("./db/pokemon.db")

	// fmt.Println(db.EmptyTableData())

	// fmt.Println(db.FillDatabase(500))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", GetAListOfPokemon).Methods("GET")

	fmt.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server started on port 8080")

}
