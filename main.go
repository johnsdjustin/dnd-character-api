/*
The main entry point to the DnD Character API.
*/
package main

import (
	"github.com/dnd-character-api/pkg/apis"
	"github.com/dnd-character-api/pkg/repositories"
)


func main(){

	characterAPIRepository := repositories.CharacterAPIRepostiory{
		URL: "https://my-json-server.typicode.com/johnsdjustin/character-sheets-db/characters",
	}

	characterAPI := apis.CharacterAPI{
		Repository: characterAPIRepository,
	}

	characterAPI.Start("localhost:8080")


}

