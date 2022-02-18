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

### Project Goals

MHs -
Create a rest API that let's users

GET all pre-generated DnD character sheets
GET all pre-generate DnD character sheets by class
GET specific pre-generated DnD character sheets by id
POST new character sheets
Data saved to local json file
Project structure follows Golang best practices
REST endpoints designed with best practices
Code structured into appropriate layers
Character sheets should be compatible with roll20

N2Hs -
Character sheet validation
ID uniqueness checks on POST
Data saved to database
Unit Tests written
Linter Integrated
Project Dockerized
Project Hosted on AWS
Automated Testing Integrated


#### Learning Goals:

Scopes
Concurrency
Data structures
Project Structure
Interfaces
JSON
