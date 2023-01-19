package seed_user_postgres_repository

import (
	"net/mail"

	"github.com/uni-school/user-microservice/libs/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), config.AppConfig.PasswordHashing.HashSalt)
	return string(bytes), err
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
