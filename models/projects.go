package models

import "github.com/hxhieu/composer-buddy/common"

// CreateProjectResult represents the result of create project API
type CreateProjectResult struct {
	Name string `json:"name"`
}

// ProjectListItem represents a project high level view
type ProjectListItem struct {
	Name   string               `json:"name"`
	Status common.ProjectStatus `json:"status"`
}
