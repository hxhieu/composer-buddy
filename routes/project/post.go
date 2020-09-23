package project

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/render"
	"github.com/hxhieu/composer-buddy/common"
	"github.com/hxhieu/composer-buddy/models"
)

// Post is creating the project
func Post(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1048576) // 1 MB

	// TODO: Go routines?

	// Get project name (actually folder name)
	name := r.FormValue("name")
	if len(name) == 0 {
		render.Status(r, 400)
		render.JSON(w, r, models.HTTPResponse{Error: "Name is required"})
		return
	}
	projectPath := filepath.Join("./", common.ProjectDir, name)

	// Check existing
	if stat, _ := os.Stat(projectPath); stat != nil {
		render.Status(r, 400)
		render.JSON(w, r, models.HTTPResponse{Error: fmt.Sprintf("The project '%s' has already existed.", name)})
		return
	}

	// Mkdir
	if err := os.MkdirAll(projectPath, os.ModePerm); err != nil {
		render.Status(r, 400)
		render.JSON(w, r, models.HTTPResponse{Error: err.Error()})
		return
	}

	processError := false

	// Get and save docker-compose.yml
	if file, _, err := r.FormFile("dockerCompose"); err != nil {
		render.Status(r, 400)
		render.JSON(w, r, models.HTTPResponse{Error: err.Error()})
		processError = true
	} else {
		if bytes, err := ioutil.ReadAll(file); err != nil {
			render.Status(r, 400)
			render.JSON(w, r, models.HTTPResponse{Error: err.Error()})
			processError = true
		} else {
			if err := ioutil.WriteFile(filepath.Join(projectPath, "docker-compose.yml"), bytes, 0755); err != nil {
				render.Status(r, 400)
				render.JSON(w, r, models.HTTPResponse{Error: err.Error()})
				processError = true
			}
		}
	}
	if processError == true {
		return
	}

	// Get and save up args
	up := r.FormValue("up")
	if err := ioutil.WriteFile(filepath.Join(projectPath, "up"), []byte(up), 0755); err != nil {
		render.Status(r, 400)
		render.JSON(w, r, models.HTTPResponse{Error: err.Error()})
		return
	}

	// Get and save down args
	down := r.FormValue("down")
	if err := ioutil.WriteFile(filepath.Join(projectPath, "down"), []byte(down), 0755); err != nil {
		render.Status(r, 400)
		render.JSON(w, r, models.HTTPResponse{Error: err.Error()})
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, models.HTTPResponse{
		Data: models.CreateProjectResult{
			Name: name,
		},
	})
}
