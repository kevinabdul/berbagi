package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"berbagi/utils/jwt"
	"berbagi/utils/password"
	"errors"
	// "strings"
	// "fmt"
)

func LoginUser(user *models.LoginUserAPI) (string ,error) {
	loginSearch := models.LoginSearchAPI{}

	res := config.Db.Table("users").Where("email = ?", user.Email).First(&loginSearch)

	if res.RowsAffected == 0 {
		return "", errors.New("No donors with corresponding email")
	}
	
	if res.Error != nil {
		return "", res.Error
	}

	if _, err := password.Check(loginSearch.Password, user.Password); err != nil {
		if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			return "", errors.New("Given password is incorrect")
		}
		return "", err
	}

	token, err := implementjwt.CreateToken(int(loginSearch.ID), "donor")

	if err != nil {
		return "", err
	}

	return token, nil

}