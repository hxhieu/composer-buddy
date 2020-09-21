package project

import (
	"net/http"

	"github.com/go-chi/render"
)

// Patch is updating the project
func Patch(w http.ResponseWriter, r *http.Request) {
	render.Status(r, 204)
}
