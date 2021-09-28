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

func GetAllProficiencies() (interface{}, int, error) {
	proficiency := []models.Proficiency{}

	tx := config.Db.Find(&proficiency)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		return proficiency, 1, nil
	}
	return nil, 0, nil
}

func DeleteProficiency(proficiencyId int) (interface{}, int, error) {
	proficiency := models.Proficiency{}

	tx := config.Db.Delete(&proficiency, proficiencyId)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		return "deleted", 1, nil
	}
	return nil, 0, nil
}

func UpdateProficiency(proficiencyId int, newProficiency *models.Proficiency) (interface{}, int, error) {
	proficiency := models.Proficiency{}

	tx := config.Db.Find(&proficiency, proficiencyId)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		update := config.Db.Model(&proficiency).Updates(newProficiency)
		if update.Error != nil {
			return nil, 0, update.Error
		}
		return proficiency, 1, nil
	}

	return nil, 0, nil
}
