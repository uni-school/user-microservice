package util

import (
	"github.com/uni-school/user-microservice/shared/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), config.AppConfig.PasswordHashing.HashSalt)
	return string(bytes), err
}
