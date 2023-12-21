package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/owenyeo/sample-react-app/backend/handlers/users"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.HandleList(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})

	}
}

func UserRoutes(r chi.Router) {
	r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
		response, _ := users.HandleList(w, req)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	r.Post("/users/login", users.LoginHandler)
	r.Post("/users/new", users.NewUserHandler)
}