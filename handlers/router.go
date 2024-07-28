package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Response struct {
	Msg  string
	Code int
}

func CreateRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CRSF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	//route adds a path
	router.Route("/api", func(r chi.Router) {
		router.Get("/healthcheck", HealthCheck)
		router.Post("/todos/create", CreateTodo)
		router.Get("/todos", GetTodos)
		router.Get("/todos/{id}", GetTodoById)
		router.Put("/todos/{id}", UpdateTodo)
		router.Delete("/todos/{id}", DeleteTodo)
	})
	return router
}
