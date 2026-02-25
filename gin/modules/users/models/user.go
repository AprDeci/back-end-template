package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"not null" json:"password"`
	NickName string `gorm:"not null" json:"nick_name"`
	Email    string `gorm:"not null;unique" json:"email"`
	gorm.Model
}

// SetPassword 加密设置密码
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPassword 验证密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
