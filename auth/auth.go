package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/universalmacro/common/singleton"
	"github.com/universalmacro/common/utils"
)

func NewSingletonJwtSigner(secret []byte) func() *JwtSigner {
	return singleton.EagerSingleton(func() *JwtSigner {
		return NewJwtSigner(secret)
	})
}

func NewJwtSigner(secret []byte) *JwtSigner {
	return &JwtSigner{
		secret: secret,
	}
}

type JwtSigner struct {
	secret []byte
}

func (j *JwtSigner) SignJwt(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.secret)
	return tokenString, err
}

func (j *JwtSigner) VerifyJwt(tokenString string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secret, nil
	})
	if token == nil {
		return nil, errors.New("unvalid token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("unvalid token")
}

type Password struct {
	Password string `json:"password" gorm:"type:CHAR(128)"`
	Salt     []byte `json:"salt"`
}

func (p *Password) SetPassword(password string) (string, []byte) {
	hashed, salt := utils.HashWithSalt(password)
	p.Password = hashed
	p.Salt = salt
	return hashed, salt
}

func (p *Password) PasswordMatching(password string) bool {
	return utils.PasswordsMatch(p.Password, password, p.Salt)
}
