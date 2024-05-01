package hash

import "golang.org/x/crypto/bcrypt"

type HashService struct{}

func NewHash() *HashService {
	return &HashService{}
}

func (h *HashService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func (h *HashService) CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
