package rest

import (
	"encoding/json"
	"net/http"
)

type SensorData struct {
	SoilMoisture int `json:"SoilMoisture,omitempty"`
	Temperature  int `json:"temperature,omitempty"`
	Humidity     int `json:"humidity,omitempty"`
}

func AddSensorData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// // Check BasicAuth login
	// if users.CheckWebAuth(r, users.UserPerm) == false {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	// Deserialize request
	var inputData SensorData
	_ = json.NewDecoder(r.Body).Decode(&inputData)

	w.WriteHeader(http.StatusCreated)
}
