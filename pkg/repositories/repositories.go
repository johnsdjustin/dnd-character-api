package repositories

import (
	"fmt"
	"log"
	"encoding/json"
	"github.com/dnd-character-api/pkg/models"
	"github.com/dnd-character-api/pkg/utils"
)

type Repository interface {
	GetCharacter(id string) models.CharacterSheet
	GetAllCharacters() []models.CharacterSheet
	CreateCharacter(data models.CharacterSheet) (bool, models.CharacterSheet)
}

type CharacterAPIRepostiory struct {
	URL string
}

func (c CharacterAPIRepostiory) GetURL() string {
	return c.URL
}

func (c CharacterAPIRepostiory) formatURL(id string) string{
	url := c.GetURL()
	return fmt.Sprintf("%s/%s", url, id)
}

func (c CharacterAPIRepostiory) GetCharacter(id string) models.CharacterSheet {
	var character models.CharacterSheet
	formattedURL := c.formatURL(id)
	body, code, err := utils.Fetch(formattedURL)

	// TODO: Return status code
	if err != nil || code >=  400 {
		log.Fatalln(err)
	}

	jErr := json.Unmarshal(body, &character)

	// TODO: Return status code
	if jErr != nil {
		log.Fatalln(jErr)
	}

	return character
}

func (c CharacterAPIRepostiory) GetAllCharacters() []models.CharacterSheet {
	var characters []models.CharacterSheet
	url := c.GetURL()

	body, code, err := utils.Fetch(url)

	// TODO: Return status code
	if err != nil || code >= 400 {
		log.Fatalln(err)
	}

	jErr := json.Unmarshal(body, &characters)

	// TODO: Return status code
	if jErr != nil {
		log.Fatalln(err)
	}

	return characters
}

func (c CharacterAPIRepostiory) CreateCharacter(data models.CharacterSheet) (bool, models.CharacterSheet){
	var newCharacter models.CharacterSheet
	url := c.GetURL()
	contentType := "application/json"

	body, code, err := utils.Create(url, contentType, data)

	// TODO: Return status code
	if err != nil || code >= 400 {
		log.Fatalln(err)
	}

	jErr := json.Unmarshal(body, &newCharacter)

	// TODO: Return status code
	if jErr != nil {
		log.Fatalln(jErr)
	}

	return true, newCharacter

}