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
// Return(s): A []byte representing the reponse body of the request in JSON, 
// 			  the HTTP status code and an error.
func Fetch(url string) ([]byte, int, error){
	resp, err := http.Get(url)
	if err != nil {
		return nil, 500, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 500, err
	}

	return body, resp.StatusCode, nil
}

// Send an HTTP Post request to create a new CharacterSheet
// Param(s): A string representing the url to create the 
// 			 CharacterSheet at, the content type being sent, and the 
// 			 CharacterSheet to be created
// Return(s): A []byte representing the reponse body of the request in JSON, 
// 			  the HTTP status code and an error.
func Create(url string, contentType string, data models.CharacterSheet) ([]byte, int, error) {

	json, err := json.Marshal(data)

	if err != nil {
		log.Fatalln(err)
		return nil, 500, err
	}

	reqBody := bytes.NewBuffer(json)
	
	resp, err := http.Post(url, contentType, reqBody)
	if err != nil {
		log.Fatalln(err)
		return nil, 500, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, 500, err
	}

	return respBody, resp.StatusCode, nil
}


type ErrorResponse struct {
	Message string `json:"message"`
}

func MapErrorResponse(code int) string {

	var errorRespString string

	// TODO: Map more status codes to strings as needed
	switch {
		case code == http.StatusNotFound:
			errorRespString = "Resource Not Found"
		case code == http.StatusBadRequest:
			errorRespString = "Bad Request"
		case code >= 500:
			errorRespString = "Server Error"
		default:
			errorRespString = "Oops! Something went wrong. Please try again!"
	}

	return errorRespString
}



func GenerateError(code int) ErrorResponse {
	errorMessage := MapErrorResponse(code)
	errorResponse := ErrorResponse{
		Message: errorMessage,
	}
	// errorRespMap := map[string]string{
	// 	"Message": errorMessage,
	// }

	return errorResponse
}

func getCharacters(c *gin.Context){
	var characters []models.CharacterSheet
	url := "https://my-json-server.typicode.com/johnsdjustin/character-sheets-db/charactersAAA"

	body, code, err := Fetch(url)
	if err != nil || code >= 400 {
		c.JSON(code, GenerateError(code))
		return
	}

	jErr := json.Unmarshal(body, &characters)
	if jErr != nil {
		c.JSON(500, GenerateError(500))
		return
	}

	c.JSON(http.StatusOK, characters)
}

func getCharactersById(c *gin.Context){
	id := c.Param("id")
	baseUrl := "https://my-json-server.typicode.com/johnsdjustin/character-sheets-db/characters"
	resourceUrl := fmt.Sprintf("%s/%s", baseUrl, id)
	var character models.CharacterSheet

	body, code, err := Fetch(resourceUrl)
	if err != nil {
		c.JSON(code, GenerateError(code))
		return
	}

	jErr := json.Unmarshal(body, &character)
	if jErr != nil {
		c.JSON(500, GenerateError(500))
		return
	}

	c.JSON(http.StatusFound, character)
}


func postCharacters(c *gin.Context){
	var newCharacter models.CharacterSheet
	url := "https://my-json-server.typicode.com/johnsdjustin/character-sheets-db/characters"
	contentType := "application/JSON"

	err := c.BindJSON(&newCharacter)
	if err != nil {
		c.JSON(500, GenerateError(500))
		return
	}

	body, code, err := Create(url, contentType, newCharacter)
	if err != nil {
		c.JSON(code, GenerateError(code))
		return
	}

	if code != http.StatusCreated {
		c.JSON(500, GenerateError(500))
		return
	}

	c.JSON(http.StatusCreated, body)
}

func main(){

	router := gin.Default(); 
	router.GET("/characters", getCharacters)
	router.GET("/characters/:id", getCharactersById)
	router.POST("characters", postCharacters)
	router.Run("localhost:8080"); 

}

