package rest

import (
	"Garden_Server/database"
	"log"
	"net/http"
)

// checkLogin validates a User type against the database and ensures that the password and email match.
func checkControllerLogin(r *http.Request) bool {
	// Pull variables from basicauth
	controller, key, ok := r.BasicAuth()
	if ok == false {
		return false
	}

	// Prepare our Database query
	stmt, err := database.DB.Prepare("SELECT count(*) FROM controllers WHERE serial = ? AND api_key = ?")
	if err != nil {
		log.Println("[checkControllerLogin - Prepare]", err)
		return false
	}
	defer stmt.Close()

	// Run our query and store the results in user variable
	var count int
	err = stmt.QueryRow(controller, key).Scan(&count)

	if err != nil {
		log.Println("[checkControllerLogin - Scan]: ", err)
		return false
	}

	if count != 1 {
		return false
	}

	return true
}
