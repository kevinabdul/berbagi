package password

import (
	"os"
	"strconv"
	bcrypt "golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	cost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func Check(dbPassword, reqPassword  string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(reqPassword))

	if err != nil {
		return false, err
	}

	return true, nil
}