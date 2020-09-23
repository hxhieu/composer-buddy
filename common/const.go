package common

const (
	// ProjectDir is the constant for the path of project files
	ProjectDir = "_projects"
)

// ProjectStatus represents the status of the given project
type ProjectStatus string

const (
	// Stop is a status of the given project
	Stop ProjectStatus = "STOP"
	// Runnning is a status of the given project
	Runnning ProjectStatus = "RUNNINg"
)
