package database

import (
	"database/sql"
	"strings"

	model "github.com/freedow10/prowarehouse-pokemon/Model"
	"github.com/freedow10/prowarehouse-pokemon/pokemon"
	"github.com/mtslzr/pokeapi-go/structs"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	dbLocation string
}

func InitDatabase(db string) Database {
	return Database{dbLocation: db}
}

func (d Database) GetAListOfPokemonFromDB(limiter int) ([]model.Pokemon, error) {
	sqliteDatabase, _ := sql.Open("sqlite3", d.dbLocation)

	rows, err := sqliteDatabase.Query(`select * from pokemon limit ?`, limiter)
	if err != nil {
		sqliteDatabase.Close()
		return nil, err
	}

	var pokemon model.Pokemon
	var pokemonList []model.Pokemon

	var tmpPMoves string
	var tmpPTypes string

	for rows.Next() {
		rows.Scan(&pokemon.Name, &pokemon.Weight, &pokemon.Height, &tmpPMoves, &tmpPTypes)
		pokemon.Moves = strings.Split(tmpPMoves, ",")
		pokemon.Types = strings.Split(tmpPTypes, ",")

		pokemonList = append(pokemonList, pokemon)

	}

	sqliteDatabase.Close()
	return pokemonList, nil
}

func (d Database) EmptyTableData() string {
	sqliteDatabase, _ := sql.Open("sqlite3", d.dbLocation)

	emptyTableSQL := `DELETE FROM pokemon;`

	statement, err := sqliteDatabase.Prepare(emptyTableSQL)
	if err != nil {
		sqliteDatabase.Close()
		return err.Error()

	}

	_, err = statement.Exec()
	if err != nil {
		sqliteDatabase.Close()
		return err.Error()
	}

	return "Table emptied"
}

func (d Database) FillDatabase(amount int) string {

	result, err := pokemon.FetchAListPokemon(amount)
	if err != nil {

		return err.Error()
	}

	modelPokemon := []model.Pokemon{}
	for _, elements := range result {

		modelPokemon = append(modelPokemon, convertStructPokemonToPokemonModel(elements))
	}

	sqliteDatabase, _ := sql.Open("sqlite3", d.dbLocation)

	insertPokemonSQL := `INSERT INTO pokemon(Name, Weight, Height,Moves,Types) VALUES (?, ?, ?, ?, ?)`

	statement, err := sqliteDatabase.Prepare(insertPokemonSQL)
	if err != nil {
		sqliteDatabase.Close()
		return err.Error()
	}

	for _, e := range modelPokemon {
		_, err = statement.Exec(e.Name, e.Weight, e.Height, e.GetPokemonMoves(), e.GetPokemonTypes())
		if err != nil {
			sqliteDatabase.Close()
			return err.Error()
		}
	}

	sqliteDatabase.Close()
	return "Database filled with records"
}

func convertStructPokemonToPokemonModel(input *structs.Pokemon) model.Pokemon {

	var tmpMoveSlice []string
	var tmpTypeSlice []string

	for _, m := range input.Moves {
		tmpMoveSlice = append(tmpMoveSlice, m.Move.Name)
	}

	for _, t := range input.Types {
		tmpTypeSlice = append(tmpTypeSlice, t.Type.Name)
	}

	return model.Pokemon{
		Name:   input.Name,
		Weight: float32(input.Weight),
		Height: float32(input.Height),
		Moves:  tmpMoveSlice,
		Types:  tmpTypeSlice,
	}
}
