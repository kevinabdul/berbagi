package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"berbagi/utils/jwt"
	"berbagi/utils/password"
	"errors"
	"strings"
)

func LoginUser(user *models.UserAPI) (string ,error) {
	role := fmt.Sprintf("%vs", strings.ToLower(user.Role))

	if role != "donators" || role != "volunteers" || role != "yayasans" || role != "personalrecipients" {
		return "", errors.New("Specified role doesnt exist")
	}

	targetUser := models.User{}

	if role == "userrecipients" {
		role = "user_recipients"
	}

	res := config.Db.Table(role).Where("email = ?", user.Email).First(&targetUser)

	if res.RowsAffected == 0 {
		return "", errors.New(fmt.Sprintf("No %v with corresponding email", role[0 : len(role) - 1])
	}
	
	if res.Error != nil {
		return "", res.Error
	}

	if _, err := password.Check(targetUser.Password, user.Password); err != nil {
		if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			return "", errors.New("Given password is incorrect")
		}
		return "", err
	}

	token, err := implementjwt.CreateToken(int(targetUser.ID), role[0 : len(role) - 1])

	if err != nil {
		return "", err
	}

	return token, nil

}package libdatabase