package rest

import (
	"Garden_Server/database"
	"encoding/json"
	"log"
	"net/http"
)

// SensorData is the structure to hold the values that we will serialze to and from the clients.
type SensorData struct {
	SoilMoisture int     `json:"soilMoisture,omitempty"`
	Temperature  float32 `json:"temperature,omitempty"`
	Humidity     float32 `json:"humidity,omitempty"`
	Time         string  `json:"time,omitempty"`
}

// AddSensorData is the REST endpoint for adding sensor data to our database.
// This will validate the controller id and key.
func AddSensorData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check BasicAuth login
	if checkControllerLogin(r) == false {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Deserialize request
	var inputData SensorData
	_ = json.NewDecoder(r.Body).Decode(&inputData)

	// Validate input
	if inputData.SoilMoisture < 1 || inputData.Humidity < 1 || inputData.Temperature < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	insertStmt, err := database.DB.Prepare("INSERT INTO sensor_data(controller_id,moisture,temperature,humidity) VALUES (?,?,?,?)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[AddSensorData - Prepare]: ", err)
		return
	}
	defer insertStmt.Close()

	controllerID, _, ok := r.BasicAuth()
	if ok == false {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[AddSensorData - getControllerId]: ", err)
		return
	}

	_, err = insertStmt.Exec(controllerID, inputData.SoilMoisture, inputData.Temperature, inputData.Humidity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[AddSensorData - Exec]: ", err)
		return
	}

	// Return a 201 for success
	w.WriteHeader(http.StatusCreated)
}

// GetSensorData is the REST endpoint for listing sensor data for a specific controller id.
// This will validate the controller id and key, though this should be moved to the users.
// @todo move validation to user accounts when completed.
func GetSensorData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check BasicAuth login
	if checkControllerLogin(r) == false {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	listStmt, err := database.DB.Prepare("SELECT moisture,temperature,humidity,time FROM sensor_data where controller_id = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[GetSensorData - Prepare]: ", err)
		return
	}

	defer listStmt.Close()

	controllerID, _, ok := r.BasicAuth()
	if ok == false {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[GetSensorData - getControllerId]: ", err)
		return
	}

	rows, err := listStmt.Query(controllerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("[GetSensorData - Query]: ", err)
		return
	}
	defer rows.Close()

	// Create a slice of SensorData to hold our results
	dataList := []SensorData{}
	count := 0
	for rows.Next() {
		count++
		var data SensorData
		err := rows.Scan(&data.SoilMoisture, &data.Temperature, &data.Humidity, &data.Time)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("[GetSensorData - Scan]: ", err)
			return
		}
		dataList = append(dataList, data)
	}
	if count == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Return our results
	json.NewEncoder(w).Encode(dataList)
}
