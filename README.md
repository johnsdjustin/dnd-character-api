### Description
A simple RESTful API written in go that will return roll20 compliant DnD character sheets. 

My goals with this project are: 

    1. To create a functioning RESTful API that returns DnD character sheets
    
    2. To learn Go

### Usage
* Clone the repo
* Start the project with go run .
* Call the API at http://localhost:8080

GET /characters

GET /characters/:id

POST /characters

POST Body:

```
{
	id string
	name string
	class string
	health int
	armorClass int
}
```

**NOTE**: Resources created with POST are not persisted. 

### Project Todo

* Update project file structure
* Return errors from CharacterAPIRepository methods
* Finish comments
* Write unit tests and automate them
* Accept Roll20 compliant character sheets
* Integrate database
* Dockerize API to make deployment simpler
* Host project
