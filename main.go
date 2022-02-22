/*
The main entry point to the DnD Character API.
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dnd-character-api/models"
	"github.com/gin-gonic/gin"
)

// Sends an HTTP GET request to fetch a resource from the given url.
// Param(s): A string representing the url to send the request to.
// Return(s): A []byte representing the reponse body in JSON, the HTTP 
// status code and an error.
func fetch(url string) ([]byte, int, error){
	resp, err := http.Get(url)
	if err != nil {
		return nil, 500, err
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 500, err
	}

	return body, resp.StatusCode, nil
}

func getCharacters(c *gin.Context){
	var characters []models.CharacterSheet
	url := "https://my-json-server.typicode.com/johnsdjustin/character-sheets-db/characters"

	body, code, err := fetch(url)
	if err != nil || code >= 500{
		log.Fatalln(err)
	}

	jErr := json.Unmarshal(body, &characters)
	if jErr != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, characters)
}

func getCharactersById(c *gin.Context){
	id := c.Param("id")
	baseUrl := "https://my-json-server.typicode.com/johnsdjustin/character-sheets-db/characters"
	finalUrl := fmt.Sprintf("%s/%s", baseUrl, id)
	var character models.CharacterSheet

	body, code, err := fetch(finalUrl)
	if err != nil || code >= 500 {
		log.Fatalln(err)
	}

	if code == 404 {
		c.JSON(code, gin.H{"Message": "Not Found"})
		return
	}

	jErr := json.Unmarshal(body, &character)
	if jErr != nil {
		log.Fatalln(jErr)
	}

	c.JSON(http.StatusFound, character)
}

func postCharacters(c *gin.Context){
	var newCharacter models.CharacterSheet
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
	router.GET("/characters/:id", getCharactersById)
	router.POST("characters", postCharacters)
	router.Run("localhost:8080"); 

}

