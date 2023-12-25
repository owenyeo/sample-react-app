package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/owenyeo/sample-react-app/backend/routes"
)

// App version
const version = "1.0.0"

type config struct {
	// App configuration type
	port int
	env  string
}

type AppStatus struct {
	// Type for the status check
	Status      string `json:"status"`
	Environment string `json:"env"`
	Version     string `json:"version"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/users", routes.UserRoutes())
	r.Route("/posts", routes.PostRoutes())

	http.ListenAndServe(":8000", r)
}
