package crypto

import (
	"election-service/configs"

	"golang.org/x/crypto/bcrypt"
)

func IsHashPasswordMatched(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string, out *string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), configs.App.Cost)
	if err != nil {
		return err
	}

	hashed := string(bytes)
	*out = hashed

	return nil
}
