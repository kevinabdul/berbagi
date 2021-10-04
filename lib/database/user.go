package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"errors"
)

func CheckUserRoleRightness(userId uint, role string) (bool, error) {
	var model interface{}

	if role == "admin" {
		model = models.Admin{}
	} else if role == "donor" {
		model = models.Donor{}
	} else if role == "volunteer" {
		model = models.Volunteer{}
	} else if role == "children" {
		model = models.Children{}
	} else if role == "foundation" {
		model = models.Foundation{}
	} else {
		return false, errors.New("invalid role")
	}

	tx := config.Db.Where("user_id = ?", userId).Find(&model)
	if tx.Error != nil {
		return false, tx.Error
	} else if tx.RowsAffected == 0 && tx.Error == nil {
		return false, nil
	}
	return true, nil
}