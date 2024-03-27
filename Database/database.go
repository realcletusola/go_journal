package database  

// import packages  
import (
	"database/sql"
	"log"

	_"github.com/lib/pq"
)


var db *sql.DB  

// journal struct represent the structure of a post  
type Journal struct{
	ID			int 	`json:"id"`
	Title		string	`json:"title"`
	Content		string	`json:"content"`
}


// InitDb initializes the database connection
func InitDb(connectionString string){
	db, err := sql.Open("postgres", connectionString)
	// check for error 
	if err != nil {
		log.Fatal(err)
	}

	// test database connection 
	err = db.Ping()
	// check for err 
	if err != nil {
		log.Println("Unable to connect to database")
		log.Fatal(err)
	}

	// if there's no error 
	log.Println("Connected to database")

	// close database connection
	defer db.Close()

}

