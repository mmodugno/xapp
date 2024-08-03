package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// HealthCheck: tests if app is working
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Msg:  "Health Check",
		Code: http.StatusOK,
	}

	jsonResponse, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}
