package libdb

import (
	"berbagi/config"
	"berbagi/models"
)

func CreateNewProficiency(newProficiency *models.Proficiency) (interface{}, int, error) {
	tx := config.Db.Create(&newProficiency)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		return newProficiency, 1, nil
	}
	return nil, 0, nil
}
