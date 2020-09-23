package project

import (
	"io/ioutil"
	"net/http"

	"github.com/hxhieu/composer-buddy/common"
	"github.com/hxhieu/composer-buddy/models"

	"github.com/go-chi/render"
)

// List is getting all projects
func List(w http.ResponseWriter, r *http.Request) {
	if files, err := ioutil.ReadDir(common.ProjectDir); err != nil {
		render.Status(r, 500)
		render.JSON(w, r, models.HTTPResponse{Error: err.Error()})
	} else {
		var projects []models.ProjectListItem
		for _, file := range files {
			if file.IsDir() {
				projects = append(projects, models.ProjectListItem{
					Name:   file.Name(),
					Status: common.Stop,
				})
			}
		}
		render.Status(r, 200)
		render.JSON(w, r, models.HTTPResponse{Data: projects})
	}
}
