package handler 


import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cletushunsu/go_journal/Database"
	"github.com/go-chi/chi/v5"

)


// handler to get all journals 
func GetAllJournals(w http.ResponseWriter, r *http.Request){
	// query database to retrieve all journals 
	rows, err := database.db.Query("SELECT * FROM journals")	
	// check for errors
	if err != nil{
		log.Println(err)
		http.Error(w, "Internal Server Error.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	journals := []database.Journal{}

	for rows.Next() {
		
		var j database.Journal 

		// scan database for all journals 
		err := rows.Scan(&j.ID, &j.Title, &j.Content)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return 
		}
		// append journals struct
		journals = append(journals, j)
	}
	// return json object 
	json.NewEncoder(w).Encode(journals)
}



// get single journal by id from database
func GetJournal(w http.ResponseWriter, r *http.Request){
	// get journal id 
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid journal id", http.StatusBadRequest)
		return 
	}

	var j database.Journal

	// query databse for journal 
	err = database.db.QueryRow("SELECT * FROM journals WHERE  id=$1", id).Scan(&j.ID, &j.Title, &j.Content)
	// check for errors 
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Journal not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	// retrun json object 
	json.NewEncoder(w).Encode(j)
} 


// create new journal in the database 
func CreateJournal(w http.ResponseWriter, r *http.Request){
	
	var j database.Journal 

	err := json.NewDecoder(r.Body).Decode(&j)
	// check for errors
	if err != nil {
		http.Error(w, "Invalid data input, please check the form and try again", http.StatusBadRequest)
		return 
	}

	_, err = database.db.Exec("INSERT INTO Journals (title, content) VALUES ($1, $2)", j.Title, j.Content)
	// check for errors
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	// return success header 
	w.WriteHeader(http.StatusCreated)

}


// updated journal 
func UpdateJournal(w http.ResponseWriter, r * http.Request){
	// get journal id 
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid journal ID", http.StatusBadRequest)
		return 
	}

	var j database.Journal

	err := json.NewDecoder(r.Body).Decode(&j)
	if err != nil {
		http.Error(w, "Invalid data input, please check form and try again", http.StatusBadRequest)
		return 
	}

	// update journal 
	_, err := database.db.Exec("UPDATE journals SET title=$1, content=$2 WHERE id=$3", j.Title, j.Content, id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	// send success header 
	w.WriteHeader(http.StatusOK)
}


// delete journal 
func DeleteJournal(w http.ResponseWriter, r *http.Request){
	// get journal id 
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil{
		http.Error(w, "Invalid journal ID", http.StatusBadRequest)
		return
	}

	// get and delete journal 
	_, err := database.db.Exec("DELETE FROM journals WHERE id=$1", id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return 
	}

	w.WriteHeader(http.StatusOK)

}
