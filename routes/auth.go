package routes

import (
	"github.com/go-chi/chi"
	"github.com/hxhieu/composer-buddy/routes/auth"
)

// AuthRoute defines all routes for /auth
func AuthRoute(r chi.Router) {
	r.Post("/login", auth.Login)
}
