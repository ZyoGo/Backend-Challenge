package core

type Auth struct {
	ID           string
	Username     string
	Email        string
	AccessToken  string
	RefreshToken string
	ExpiredAt    int64
}
