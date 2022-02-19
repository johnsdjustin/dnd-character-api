package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func getCharacters(c *gin.Context){

	var characters []CharacterSheet

	resp, err := http.Get("https://my-json-server.typicode.com/johnsdjustin/character-sheets-db/characters")

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	jErr := json.Unmarshal(body, &characters)

	if jErr != nil {
		log.Fatalln(jErr)
	}

	c.JSON(http.StatusOK, characters)
}

func getCharactersById(c *gin.Context){
	id := c.Param("id")

	var characters []CharacterSheet

	resp, err := http.Get("https://my-json-server.typicode.com/johnsdjustin/character-sheets-db/characters")

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	jErr := json.Unmarshal(body, &characters)

	if jErr != nil {
		log.Fatalln(jErr)
	}

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

	jsonBody, err :=  json.Marshal(newCharacter)

	if err !=  nil {
		log.Fatalln(err)
	}

	postBody := bytes.NewBuffer(jsonBody)

	resp, err := http.Post("https://my-json-server.typicode.com/johnsdjustin/character-sheets-db/characters", "application/JSON", postBody)
	
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	if resp.StatusCode != http.StatusCreated {
		c.JSON(http.StatusCreated, gin.H{"message": "Character creation failed. Please try again"})
		return
	}

	c.JSON(http.StatusCreated, newCharacter)
}

func main(){

	router := gin.Default(); 
	router.GET("/characters", getCharacters)
	// router.GET("/characters/:id", getCharactersById)
	router.POST("characters", postCharacters)
	router.Run("localhost:8080"); 

}

