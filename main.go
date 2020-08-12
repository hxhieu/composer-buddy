package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/hxhieu/composer-buddy/routes"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Routes
	r.Route("/api/auth", routes.AuthRoute)

	// Port
	port := os.Getenv("COMPOSER_BUDDY_PORT")
	if len(port) == 0 {
		port = "5880"
	}

	// Start
	fmt.Printf("Server listening at port :%s...Ctrl+C to stop.\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		fmt.Printf("Error %s\n", err)
	}
}
