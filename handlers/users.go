package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"x-app-go/services"
)

// PostUser : creates a new user
func PostUser(w http.ResponseWriter, r *http.Request) {
	var user services.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ResponseBuilder(w, Response{
			Msg:  err.Error(),
			Code: http.StatusBadRequest,
		})
		return
	}

	err = validateNewUser(w, user)
	if err != nil {
		ResponseBuilder(w, Response{
			Msg:  err.Error(),
			Code: http.StatusBadRequest,
		})
		return
	}

	err = user.InsertUser(user)

	if err != nil {
		res := Response{
			Msg:  fmt.Sprintf("cannot create user: %s", err.Error()),
			Code: http.StatusInternalServerError,
		}
		ResponseBuilder(w, res)
		return
	}

	res := Response{
		Msg:  fmt.Sprintf("created user with username: %s", user.Username),
		Code: http.StatusCreated,
	}
	ResponseBuilder(w, res)

	return
}

func ResponseBuilder(w http.ResponseWriter, r Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	json.NewEncoder(w).Encode(r)
}

func validateNewUser(w http.ResponseWriter, user services.User) error {
	switch user.Username {
	case "":
		return fmt.Errorf("empty username")
	case "error":
		return fmt.Errorf("unexpected error")
	}

	return nil
}
