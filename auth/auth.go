package auth

import "github.com/universalmacro/common/utils"

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
