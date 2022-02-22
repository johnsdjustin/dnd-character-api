/*
The models package houses json types and other types reprenting data used 
throughout the DnD Character API.
*/
package models

type CharacterSheet struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Class string `json:"class"`
	Health int `json:"health"`
	ArmorClass int `json:"armorClass"`
}

type CharacterSheets struct {
	Characters []CharacterSheet `json:"characters"`
}

