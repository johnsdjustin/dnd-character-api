/*
types.go houses the types used throughout the DnD Character API
*/

package main

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

