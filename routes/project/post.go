package project

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

// Post is creating the project
func Post(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(50000)
	fmt.Println(r.FormValue("name"))
	fmt.Println(r.FormFile("dockerCompose"))
	render.Status(r, 204)
}
