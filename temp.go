package main
/**
* This is a swamp of code I want to keep... but at a distance. 
*/

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

	// var characters = []CharacterSheet {
	// 	{ID: "1", Name: "Bilbo", Class: "Burglar", Health: 10, ArmorClass: 10},
	// 	{ID: "2", Name: "Thorin", Class: "Warrior", Health: 20, ArmorClass: 20},
	// 	{ID: "3", Name: "Gandalf", Class: "Wizard", Health: 15, ArmorClass: 10},
	// }

	// sheets.Characters = append(sheets.Characters, CharacterSheet{ID: "3", Name: "Gandalf", Class: "Wizard", Health: 15, ArmorClass: 10})

	// for _, character := range characters {
	// 	sheets.AddCharacter(character)
	// 	sheets.Count++
	// }

	// type CharacterSheets struct {
// 	Count int `json:"count"`
// 	Characters []CharacterSheet `json:"characters"`
// }

// var sheets = CharacterSheets{Count: 0, Characters: make([]CharacterSheet, 0)}

// func (sheets *CharacterSheets) AddCharacter(character CharacterSheet) []CharacterSheet{
// 	sheets.Characters = append(sheets.Characters, character)
// 	return sheets.Characters
// }

// func fileDemo() {
	// Open a file and write new content back to it. 

	// data, rdErr := os.ReadFile("./test.txt")

	// if rdErr != nil {
	// 	fmt.Println(rdErr)
	// }

	// strData := string(data)
	// addData := strData + " " + "Here is some new data"

	// newData := []byte(addData)

	// wtErr := os.WriteFile("newTest.txt", newData, 0777)

	// if wtErr != nil {
	// 	fmt.Println(wtErr)
	// }

// 	f, err := os.OpenFile("./newTest.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	if _, wtErr := f.Write([]byte("New Data\n")); err != nil {
// 		fmt.Println(wtErr)
// 		f.Close()
// 	}

// 	if err := f.Close(); err != nil {
// 		fmt.Println(err)
// 	}
// }

// func jsonDemo() {

// 	type Author struct {
// 		Sales int `json:"book_sales"`
// 		Age int `json:"age"`
// 		Developer bool `json:"is_developer"`
// 	}

// 	type Book struct {
// 		Title string `json:"title"`
// 		Author Author `json:"author"` 
// 	}

// 	type SensorReading struct {
// 		Name string `json:"name"`
// 		Capacity int `json:"capacity"`
// 		Time string `json:"time"`
// 	}
	
// 	author := Author{Sales: 100, Age: 24, Developer: true}
// 	book := Book{Title: "Book", Author: author}

// 	// byteArray, err := json.Marshal(book) 
// 	byteArray, err := json.MarshalIndent(book, "", " ")
	
// 	if err != nil {
// 		fmt.Println(err)
// 	} 

// 	// fmt.Printf("%+v\n",book)

// 	fmt.Println(string(byteArray))

// 	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`

// 	// var reading SensorReading
// 	var reading map[string]interface{}

// 	umErr := json.Unmarshal([]byte(jsonString), &reading)

// 	if umErr != nil {
// 		fmt.Println(umErr)
// 	}

// 	fmt.Println(reading)

// 	os.WriteFile("characters.json", byteArray, 0644)

// }