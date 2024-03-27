package main 

import (
	"log"
	"net/http"

	"github.com/cletushunsu/go_journal/Database"
	"github.com/cletushunsu/go_journal/Router"
)


// main function 
func main() {
	// create router instance
	r := routes.NewRouter()

	// define database connection string 
	connectionString := 

	// create database instance 
	conn, err = database.InitDb(connectionString)
	// check for error 
	if err != nil {
		log.Panic(err)
	}

	// start server on port 8080 
	log.Println("Server listening on :8080")
	http.ListenAndServe(":8080", r)

}