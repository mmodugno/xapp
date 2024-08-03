package handlers

import (
	"encoding/json"
	"net/http"
	"x-app-go/internal/core/services/timeline"
)

func GetTimeline(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	var timeline timeline.Timeline

	if id == "" {
		ResponseBuilder(w, Response{
			Msg:  "please provide an ID",
			Code: http.StatusBadRequest,
		})
		return
	}
	tm, err := timeline.GetTimeline(id)
	if err != nil {
		ResponseBuilder(w, Response{
			Msg:  err.Error(),
			Code: http.StatusNotFound,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tm)
	return
}
