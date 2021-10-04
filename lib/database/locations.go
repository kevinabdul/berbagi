package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"database/sql"
	"errors"
	"fmt"

	"gorm.io/gorm"
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

	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		return address, int(tx.RowsAffected), nil
	}

	return nil, 0, nil
}

func GetRequestByAddressIdResolve(addressId uint, resolved string) ([]models.Request, int, error) {
	var request []models.Request

	var tx *gorm.DB
	if resolved == "no" {
		tx = config.Db.Where("address_id = ? AND resolved = false", addressId).Find(&request)
	} else if resolved == "yes" {
		tx = config.Db.Where("address_id = ? AND resolved = true", addressId).Find(&request)
	} else {
		tx = config.Db.Where("address_id = ?", addressId).Find(&request)
	}

	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		return request, int(tx.RowsAffected), nil
	}

	return nil, 0, nil
}

func GetUserByAddressIdRole(addressId uint, role string) (models.UserProfile, int, error) {
	var res models.UserProfile

	var tx *gorm.DB
	if role == "admin" {
		user := models.Admin{}
		tx = config.Db.Where("address_id = ?", addressId).Find(&user)
		res.UserID = user.UserID
		res.Role = "admin"
		res.Address = user.Address.Name
		res.City = user.Address.City.Name
		res.Province = user.Address.Province.Name
	} else if role == "donor" {
		user := models.Donor{}
		tx = config.Db.Where("address_id = ?", addressId).Find(&user)
		res.UserID = user.UserID
		res.Name = user.User.Name
		res.Role = "donor"
		res.Address = user.Address.Name
		res.City = user.Address.City.Name
		res.Province = user.Address.Province.Name
	} else if role == "volunteer" {
		user := models.Volunteer{}
		tx = config.Db.Where("address_id = ?", addressId).Find(&user)
		res.UserID = user.UserID
		res.Name = user.User.Name
		res.Role = "volunteer"
		res.Address = user.Address.Name
		res.City = user.Address.City.Name
		res.Province = user.Address.Province.Name
	} else if role == "children" {
		user := models.Children{}
		tx = config.Db.Where("address_id = ?", addressId).Find(&user)
		res.UserID = user.UserID
		res.Name = user.User.Name
		res.Role = "children"
		res.Address = user.Address.Name
		res.City = user.Address.City.Name
		res.Province = user.Address.Province.Name
	} else if role == "foundation" {
		user := models.Foundation{}
		tx = config.Db.Where("address_id = ?", addressId).Find(&user)
		res.UserID = user.UserID
		res.Name = user.User.Name
		res.Role = "foundation"
		res.Address = user.Address.Name
		res.City = user.Address.City.Name
		res.Province = user.Address.Province.Name
	} else if role == "" {
		user := models.User{}
		tx = config.Db.Where("address_id = ?", addressId).Find(&user)
		res.UserID = user.ID
		res.Name = user.Name
		res.Role = user.Role
	} else {
		return models.UserProfile{}, 0, errors.New("invalid role")
	}

	if tx.Error != nil {
		return models.UserProfile{}, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		return res, int(tx.RowsAffected), nil
	}

	return models.UserProfile{}, 0, nil
}
