package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"errors"
)

func CheckUserRoleRightness(userId uint, role string) (bool, error) {
	if role == "admin" {
		model := models.Admin{}
		tx := config.Db.Where("user_id = ?", userId).Find(&model)
		if tx.Error != nil {
			return false, tx.Error
		} else if tx.RowsAffected == 0 && tx.Error == nil {
			return false, nil
		}
		return true, nil
	} else if role == "donor" {
		model := models.Donor{}
		tx := config.Db.Where("user_id = ?", userId).Find(&model)
		if tx.Error != nil {
			return false, tx.Error
		} else if tx.RowsAffected == 0 && tx.Error == nil {
			return false, nil
		}
		return true, nil
	} else if role == "volunteer" {
		model := models.Volunteer{}
		tx := config.Db.Where("user_id = ?", userId).Find(&model)
		if tx.Error != nil {
			return false, tx.Error
		} else if tx.RowsAffected == 0 && tx.Error == nil {
			return false, nil
		}
		return true, nil
	} else if role == "children" {
		model := models.Children{}
		tx := config.Db.Where("user_id = ?", userId).Find(&model)
		if tx.Error != nil {
			return false, tx.Error
		} else if tx.RowsAffected == 0 && tx.Error == nil {
			return false, nil
		}
		return true, nil
	} else if role == "foundation" {
		model := models.Foundation{}
		tx := config.Db.Where("user_id = ?", userId).Find(&model)
		if tx.Error != nil {
			return false, tx.Error
		} else if tx.RowsAffected == 0 && tx.Error == nil {
			return false, nil
		}
		return true, nil
	}
	return false, errors.New("invalid role")
}
