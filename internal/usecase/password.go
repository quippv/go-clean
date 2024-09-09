package usecase

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	saltLength = 16
	keyLength  = 32
)

type PasswordUseCase struct{}

func NewPasswordUseCase() *PasswordUseCase {
	return &PasswordUseCase{}
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func (p *PasswordUseCase) HashPassword(password string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}

	hashedPassword := argon2.Key([]byte(password), salt, 1, 64*1024, 4, keyLength)
	saltEncoded := base64.RawStdEncoding.EncodeToString(salt)
	hashEncoded := base64.RawStdEncoding.EncodeToString(hashedPassword)

	return fmt.Sprintf("%s$%s", saltEncoded, hashEncoded), nil
}
