package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (user *User) HashPassword(pass string) error {
	passBytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		return err
	}
	user.Password = string(passBytes)
	return nil
}

func (user *User) CheckPassword(givenPass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(givenPass)); err != nil {
		return err
	}
	return nil
}
