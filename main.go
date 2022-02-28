/*
The main entry point to the DnD Character API.
*/
package main

import (
	"net/http"
	"github.com/dnd-character-api/pkg/models"
	"github.com/dnd-character-api/pkg/repositories"
	"github.com/dnd-character-api/pkg/utils"
	"github.com/gin-gonic/gin"
)


type CharacterAPI struct {
	Repository repositories.Repository
}

func (API CharacterAPI) postCharacters(c *gin.Context){
	var newCharacter models.CharacterSheet
	repository := API.Repository

	err := c.BindJSON(&newCharacter)
	if err != nil {
		c.JSON(500, utils.GenerateError(500))
		return
	}

	created, createdCharacter := repository.CreateCharacter(newCharacter)

	if !created {
		c.JSON(400, utils.GenerateError(400))
		return
	}

	c.JSON(http.StatusCreated, createdCharacter)
}

func (API CharacterAPI) getCharacters(c *gin.Context){

	repository := API.Repository

	characters := repository.GetAllCharacters()

	c.JSON(http.StatusOK, characters)
}

func (API CharacterAPI) getCharactersById(c *gin.Context){
	id := c.Param("id")
	repository := API.Repository

	character := repository.GetCharacter(id)

	c.JSON(http.StatusFound, character)
}

func (API CharacterAPI) Start(address string) {
	router := gin.Default(); 
	router.GET("/characters", API.getCharacters)
	router.GET("/characters/:id", API.getCharactersById)
	router.POST("characters", API.postCharacters)
	router.Run(address); 
}



func main(){

	characterAPIRepository := repositories.CharacterAPIRepostiory{
		URL: "https://my-json-server.typicode.com/johnsdjustin/character-sheets-db/characters",
	}

	characterAPI := CharacterAPI{
		Repository: characterAPIRepository,
	}

	characterAPI.Start("localhost:8080")


}

