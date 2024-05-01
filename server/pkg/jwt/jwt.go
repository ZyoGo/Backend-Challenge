package jwt

import (
	"time"

	"github.com/ZyoGo/Backend-Challange/config"
	"github.com/ZyoGo/Backend-Challange/pkg/derrors"
	"github.com/golang-jwt/jwt/v5"
)

// Bad, this should be not here. but at config and use Asymmetric JWT
var jwtKey []byte = []byte("0338ba1062563db29fe46cbd25bc3c52e2b996f60c26e8b421ac9b9fbd108fb6")

const InternalErrMsg string = "Internal Server Error"

type CustomClaimsJWT struct {
	Email  string `json:"email"`
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

type TokenJWT struct {
	AccessToken  string
	RefreshToken string
	ExpiredAt    int64
}

type AuthGuardJWT struct {
	Email  string `json:"email"`
	UserId string `json:"user_id"`
}

type jwtBusiness struct{}

func NewBusiness() *jwtBusiness {
	return &jwtBusiness{}
}

func (j *jwtBusiness) GenerateTokenJWT(email, userId string) (response TokenJWT, err error) {
	expiredAt := time.Now().Add(60 * time.Minute)
	claimsAccess := CustomClaimsJWT{
		Email:  email,
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    config.GetConfig().App.Name,
		},
	}

	claimsRefresh := claimsAccess
	claimsRefresh.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))

	tokenAcc := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccess)
	accessToken, err := tokenAcc.SignedString(jwtKey)
	if err != nil {
		derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, InternalErrMsg)
	}

	tokenRef := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	refreshToken, err := tokenRef.SignedString(jwtKey)
	if err != nil {
		derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, InternalErrMsg)
	}

	response = TokenJWT{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    expiredAt.Unix(),
	}

	return response, nil
}

func (j *jwtBusiness) ParseAndVerifyJWT(jwtToken string) (AuthGuardJWT, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &CustomClaimsJWT{}, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return AuthGuardJWT{}, derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, InternalErrMsg)
	}

	claims, ok := token.Claims.(*CustomClaimsJWT)
	if ok && token.Valid {
		return AuthGuardJWT{
			Email:  claims.Email,
			UserId: claims.UserId,
		}, nil
	}

	return AuthGuardJWT{}, err
}
