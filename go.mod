module github.com/freedow10/prowarehouse-pokemon

replace github.com/freedow10/prowarehouse-pokemon/pokemon => ../pokemon

replace github.com/freedow10/prowarehouse-pokemon/Model => ../Model

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/mattn/go-sqlite3 v1.14.8
	github.com/mtslzr/pokeapi-go v1.4.0
)
