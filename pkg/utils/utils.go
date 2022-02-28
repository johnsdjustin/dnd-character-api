package utils

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"github.com/dnd-character-api/pkg/models"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func mapErrorResponse(code int) string {

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
	errorMessage := mapErrorResponse(code)
	errorResponse := ErrorResponse{
		Message: errorMessage,
	}

	return errorResponse
}

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
		return nil, 500, err
	}

	reqBody := bytes.NewBuffer(json)
	
	resp, err := http.Post(url, contentType, reqBody)
	if err != nil {
		return nil, 500, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 500, err
	}

	return respBody, resp.StatusCode, nil
}