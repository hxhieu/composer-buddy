package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/hxhieu/composer-buddy/routes"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

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
