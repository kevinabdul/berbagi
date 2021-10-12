package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"berbagi/utils/jwt"
	"berbagi/utils/password"
	"errors"
	"fmt"
	// "strings"
)

func LoginUser(user models.LoginUserAPI) (string ,error) {
	loginSearch := models.LoginSearchAPI{}

	res := config.Db.Table("users").Select("users.id, users.email, users.password, users.role_id, roles.name as role_name").
	Joins("left join roles on roles.id = users.role_id").Where("email = ?", user.Email).First(&loginSearch)

	if res.RowsAffected == 0 {
		return "", errors.New("No user with corresponding email")
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

	fmt.Println(loginSearch)

	token, err := implementjwt.CreateToken(int(loginSearch.ID), loginSearch.RoleName)

	if err != nil {
		return "", err
	}

	return token, nil

}