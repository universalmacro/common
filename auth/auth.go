package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/universalmacro/common/config"
	"github.com/universalmacro/common/utils"
)

var secret []byte

func init() {
	secret = []byte(config.GetString("jwt.secret"))
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

func SignJwt(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	return tokenString, err
}

func VerifyJwt(tokenString string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if token == nil {
		return nil, errors.New("unvalid token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("unvalid token")
}
