package service

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/truetandem/e-QIP-prototype/api/model"
)

var (
	JwtSecret         = []byte("more secrets!")
	Issuer            = "eqip"
	BasicAuthAudience = "Basic"
	Expiration        = time.Minute * 20
)

type JwtTokenService struct {
	secret []byte
}

func (j *JwtTokenService) New(account model.Account) (string, time.Time, error) {
	expiresAt := time.Now().Add(Expiration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        string(account.ID),
		Issuer:    Issuer,
		Audience:  BasicAuthAudience,
		ExpiresAt: expiresAt.Unix(),
	})

	signedToken, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", time.Time{}, err
	}

	return signedToken, expiresAt, nil
}

// ValidJwtToken parses a token and determines if the token is valid
func (t *JwtTokenService) Valid(rawToken string) (bool, error) {
	token, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return JwtSecret, nil
	})

	// Invalid token
	if err != nil {
		return false, err
	}

	// Everything is good
	return token.Valid, err
}

func NewJwtTokenService() *JwtTokenService {
	return &JwtTokenService{JwtSecret}
}
