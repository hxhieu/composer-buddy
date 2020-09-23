package routes

import (
	"github.com/go-chi/chi"
	"github.com/hxhieu/composer-buddy/routes/project"
)

// ProjectRoute defines all routes for /project
func ProjectRoute(r chi.Router) {
	r.Post("/", project.Post)
	r.Patch("/", project.Patch)
	r.Get("/", project.List)
}
