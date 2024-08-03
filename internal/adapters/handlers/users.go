package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"x-app-go/internal/core/services/users"
)

func ResponseBuilder(w http.ResponseWriter, r Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	json.NewEncoder(w).Encode(r)
}

// PostUser : creates a new user
func PostUser(w http.ResponseWriter, r *http.Request) {
	var user users.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ResponseBuilder(w, Response{
			Msg:  err.Error(),
			Code: http.StatusBadRequest,
		})
		return
	}

	err = validateNewUser(user)
	if err != nil {
		ResponseBuilder(w, Response{
			Msg:  err.Error(),
			Code: http.StatusBadRequest,
		})
		return
	}

	err = user.InsertUser(user.Username)

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

func validateNewUser(user users.User) error {
	switch user.Username {
	case "":
		return fmt.Errorf("empty username")
	case "error":
		return fmt.Errorf("unexpected error")
	}

	return nil
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	var user users.User
	id := chi.URLParam(r, "id")

	if id == "" {
		ResponseBuilder(w, Response{
			Msg:  "please provide an ID",
			Code: http.StatusBadRequest,
		})
		return
	}
	user, err := user.GetUserByID(id)
	if err != nil {
		ResponseBuilder(w, Response{
			Msg:  err.Error(),
			Code: http.StatusNotFound,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	return
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user users.User
	id := r.Header.Get("id")

	if id == "" {
		ResponseBuilder(w, Response{
			Msg:  "please provide an ID",
			Code: http.StatusBadRequest,
		})
		return
	}
	err := user.Delete(id)
	if err != nil {
		ResponseBuilder(w, Response{
			Msg:  err.Error(),
			Code: http.StatusInternalServerError,
		})
		return
	}
	res := Response{
		Msg:  "Successfully deleted",
		Code: http.StatusOK,
	}
	ResponseBuilder(w, res)

	return
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	var user users.User
	//Get ID from user, and username of the user to follow
	id := r.Header.Get("id")
	username := chi.URLParam(r, "username")

	if username == "" {
		ResponseBuilder(w, Response{
			Msg:  "please provide an username",
			Code: http.StatusBadRequest,
		})
		return
	}

	err := user.FollowUser(id, username)
	if err != nil {
		ResponseBuilder(w, Response{
			Msg:  err.Error(),
			Code: http.StatusInternalServerError,
		})
		return
	}
	res := Response{
		Msg:  "Successfully followed",
		Code: http.StatusOK,
	}
	ResponseBuilder(w, res)

	return
}
