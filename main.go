package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	model "github.com/freedow10/prowarehouse-pokemon/Model"
	"github.com/freedow10/prowarehouse-pokemon/database"
	"github.com/gorilla/mux"
)

type record struct {
	TotalRecords int                `json:"Count"`
	Page         string             `json:"Page"`
	Results      []pokemonViewModel `json:"Results"`
}

type pokemonViewModel struct {
	Name   string   `json:"name"`
	Weight float32  `json:"weight"`
	Height float32  `json:"height"`
	Moves  []string `json:"moves"`
	Types  []string `json:"types"`
}

func GetAListOfPokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()

	pagenr := 1
	pageLimiter := 10

	if len(q.Get("page")) > 0 {
		var err error
		pagenr, err = strconv.Atoi(q.Get("page"))
		if err != nil {
			w.Write([]byte(string("Error on converting page number")))
			return
		}

		if pagenr <= 0 {
			pagenr = 1
		}
	}

	if len(q.Get("limit")) > 0 {
		var err error
		pageLimiter, err = strconv.Atoi(q.Get("limit"))
		if err != nil {
			w.Write([]byte(string("Error on converting limit value")))
			return
		}

		if pageLimiter <= 0 {
			pageLimiter = 10
		}
	}

	db := database.InitDatabase("./db/pokemon.db")

	dbResults, count, err := db.GetAListOfPokemonFromDB(pageLimiter, pagenr)
	if err != nil {
		fmt.Println(err)
		return
	}

	pokemonList := []pokemonViewModel{}

	for _, elements := range dbResults {
		pokemonList = append(pokemonList, convertPokemonModelToViewModel(elements))
	}

	currentPage := 1
	if pagenr > 0 {
		currentPage = pagenr
	}

	pageString := "Page " + strconv.Itoa(currentPage) + " of " + strconv.Itoa(int(math.Ceil(float64(count)/float64(pageLimiter))))

	results := record{
		TotalRecords: count,
		Page:         pageString,
		Results:      pokemonList,
	}

	jsonPokemon, _ := json.Marshal(results)
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
