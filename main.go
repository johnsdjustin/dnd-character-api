package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

var characters []CharacterSheet

func getCharacters(c *gin.Context){
	c.JSON(http.StatusOK, characters)
}

func getCharactersById(c *gin.Context){
	var characters = characters
	id := c.Param("id")

	for _, character := range characters{
		if id == character.ID {
			c.JSON(http.StatusOK, character)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "character not found"})
}

func postCharacters(c *gin.Context){
	var newCharacter CharacterSheet
	
	err := c.BindJSON(&newCharacter)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Interal failure - please try again"})
		return
	}

	characters = append(characters, newCharacter)

	c.JSON(http.StatusCreated, newCharacter)

}

func initializeData(){

	fmt.Println("Fetching Data...")

	resp, hErr := http.Get("https://my-json-server.typicode.com/johnsdjustin/character-sheets-db/characters")

	if hErr != nil {
		log.Fatalln(hErr)
	}

	body, iErr := ioutil.ReadAll(resp.Body)

	if iErr != nil {
		log.Fatalln(iErr)
	}

	jErr := json.Unmarshal(body, &characters)

	if jErr != nil {
		log.Fatalln(jErr)
	} else {
		fmt.Println("Fetch Complete!")
	}
}

func main(){

	initializeData()

	router := gin.Default(); 
	router.GET("/characters", getCharacters)
	router.GET("/characters/:id", getCharactersById)
	router.POST("characters", postCharacters)
	router.Run("localhost:8080"); 

}

