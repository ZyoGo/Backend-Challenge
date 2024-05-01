package core

type Hash interface {
	HashPassword(password string) (string, error)
	CompareHashPassword(password, hash string) bool
}
