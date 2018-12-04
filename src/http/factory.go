package http

//Exports provides all a model for exposing all the services ready to be used.
type Exports struct {
	Client Client
}

//PublicInstances holds all instances of http client ready to be used.
var PublicInstances *Exports
