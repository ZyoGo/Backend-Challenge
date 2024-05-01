package request

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
