package Models

import "github.com/jinzhu/gorm"

type UserPet struct {
	gorm.Model
	User User
	Pet  Pet
}
