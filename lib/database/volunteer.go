package libdb

import (
	"berbagi/config"
	"berbagi/models"
)

func ListVolunteers() (interface{}, int, error) {
	volunteerRes := []models.VolunteerAPI{}

	query := `SELECT volunteers.id, users.name, users.email 
				FROM volunteers
				INNER JOIN users ON volunteers.user_id = users.id;`

	tx := config.Db.Raw(query).Scan(&volunteerRes)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		return volunteerRes, 1, nil
	}
	return nil, 0, nil
}
