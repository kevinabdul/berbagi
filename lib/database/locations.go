package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"database/sql"
	"fmt"
)

func GetAddressLatLonByUserId(id uint, role string) (models.LocationPointResponseAPI, error) {
	var point models.LocationPointResponseAPI

	var table string
	if role == "donor" {
		table = "donors"
	} else if role == "volunteer" {
		table = "volunteers"
	}

	tx := config.Db.Table("addresses").Select(
		"addresses.latitude, addresses.longitude").Joins(
		fmt.Sprintf("JOIN %s ON addresses.id = %s.address_id", table, table)).Joins(
		fmt.Sprintf("JOIN users ON users.id = %s.user_id", table)).Where(
		"users.id = ?", id).First(&point)

	if tx.Error != nil {
		return models.LocationPointResponseAPI{}, tx.Error
	}

	return point, nil
}

func GetAllNearestAddressId(lat, lon, _range float64) ([]models.NearestAddressIdResponseAPI, int, error) {
	var address []models.NearestAddressIdResponseAPI

	// query := `SELECT id, distance
	// 		  From (
	// 			Select
	// 			( 6371 * acos( cos( radians( @lat ) )
	// 		  	* cos( radians( latitude ) )
	// 		  	* cos( radians( longitude ) - radians( @lon ) )
	// 			+ sin( radians( @lat ) )
	// 		  	* sin(radians( latitude ) ) ) ) distance
	// 	  	  	From addresses )
	// 		  Where distance < @range
	// 		  ORDER BY distance`

	// BACKUP QUERY
	query := `SELECT id, (
		6371 * acos( cos( radians( @lat ) )
		* cos( radians( latitude ) )
		* cos( radians( longitude ) - radians( @lon ) )
		+ sin( radians( @lat ) )
		* sin( radians( latitude ) ) )) AS distance
	FROM addresses
	WHERE latitude<>''
		AND longitude<>''
	HAVING distance < @range
	ORDER BY distance asc`

	tx := config.Db.Raw(query,
		sql.Named("lat", lat),
		sql.Named("lon", lon),
		sql.Named("range", _range)).Scan(&address) // All saved or per row?

	// tx := config.Db.Where(`
	// 				(latitude BETWEEN ? AND ?)
	// 				AND
	// 				(longitude BETWEEN ? AND ?)`,lat1, lat2, lon1, lon2).Find(&address)

	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		return address, int(tx.RowsAffected), nil
	}

	return nil, 0, nil
}

// SELECT id, distance
// From (Select ( 6371 * acos( cos( radians(37) )
//           * cos( radians( latitude ) )
//           * cos( radians( Longitude ) - radians(-122) ) +
//              sin( radians(37) )
//           * sin(radians(latitude)) ) ) distance
//       From DriverLocationHistory)z
// Where distance < 5
// ORDER BY distance
// type user interface {
// 	model() interface{}
// }

// type roleDonor string
// func (r *roleDonor) model() interface{} {return models.Donor{}}

// type roleVolunteer string
// func (r *roleVolunteer) model() interface{} {return models.Volunteer{}}

// func GetAddressIdByUserIdRole(userId int, role string) (int, error) {
// 	var r user
// 	if role == "donor" {
// 		r = new(roleDonor)
// 	} else if role == "volunteer" {
// 		r = new(roleVolunteer)
// 	}

// 	model := r.model()

// 	tx := config.Db.First(&model, userId)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	addressId := model.AddressID
// }
