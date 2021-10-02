package models

type ResponseOK struct {
	Status 	string		`json:"status"`
	Message string		`json:"message"`
	Data 	interface{}	`json:"data"`
}

type ResponseNotOK struct {
	Status 	string		`json:"status"`
	Message string		`json:"message"`
}