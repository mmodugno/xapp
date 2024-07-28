package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"x-app-go/services"
)

// HealthCheck: to test if app is working
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

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo services.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
	}

	err = todo.InsertTodo(todo)
	if err != nil {
		res := Response{
			Msg:  "Cannot create todo",
			Code: http.StatusInternalServerError,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := Response{
		Msg:  "Created Successfully",
		Code: http.StatusCreated,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)

	return
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todo services.Todo
	todos, err := todo.GetAllTodos()
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	var todo services.Todo
	id := chi.URLParam(r, "id")

	todo, err := todo.GetTodoById(id)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	var todo services.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = todo.UpdateTodo(id, todo)
	if err != nil {
		res := Response{
			Msg:  fmt.Sprintf("Cannot update record: %s", err),
			Code: http.StatusNotModified,
		}
		json.NewEncoder(w).Encode(res)
		w.WriteHeader(res.Code)
		return
	}

	res := Response{
		Msg:  "Updated successfully",
		Code: http.StatusOK,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	var todo services.Todo
	id := chi.URLParam(r, "id")

	err := todo.DeleteTodo(id)
	if err != nil {
		res := Response{
			Msg:  "Error",
			Code: 304,
		}
		json.NewEncoder(w).Encode(res)
		w.WriteHeader(res.Code)
		return
	}

	res := Response{
		Msg:  "Successfully deleted",
		Code: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
}
