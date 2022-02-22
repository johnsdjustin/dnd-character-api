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

Character Sheets Roll20 compliant

Character sheet validation

ID uniqueness checks on POST

Data saved to database

Automated Tests

Linter Integrated

Project Dockerized

Project Hosted on AWS



cockroach sql --url "postgresql://justin:Mhu1glYlSHigorqCc3iiVA@free-tier14.aws-us-east-1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full&options=--cluster%3Dfunky-beaver-424&sslrootcert=$HOME/.postgresql/root.crt"