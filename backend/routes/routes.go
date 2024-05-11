package routes

import (

	"github.com/go-chi/chi/v5"
	"github.com/owenyeo/sample-react-app/backend/handlers/posts"
	"github.com/owenyeo/sample-react-app/backend/handlers/users"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/users", users.HandleList)
	}
}

func UserRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/login", users.LoginHandler)
	}
}

func PostRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", posts.HandleListPosts)
		r.Post("/new", posts.HandleAddPost)
	}
}
