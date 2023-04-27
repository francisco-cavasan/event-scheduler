package Models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func (u *User) SaveUser() (*User, error) {

	var err error
	var db gorm.DB

	err = db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
