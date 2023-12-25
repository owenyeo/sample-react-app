package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/owenyeo/sample-react-app/backend/handlers/posts"
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

func UserRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/login", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.LoginHandler(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})

		r.Post("/new", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.NewUserHandler(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
	}
}

func PostRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, req *http.Request) {
			response, err := posts.HandleListPosts(w, req)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})

		r.Post("/new", func(w http.ResponseWriter, req *http.Request) {
			response, err := posts.HandleAddPost(w, req)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			fmt.Println("Successfully added post")

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
	}
}
