package response 

import (
	"berbagi/models"
)

func Create(status, message string, data interface{}) interface{}{

	if status == "success" {
		return models.ResponseOK {
			Status: "success",
			Message: message,
			Data: data}
	}

	return models.ResponseNotOK {
		Status: "failed",
		Message: message}
}