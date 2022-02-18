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

