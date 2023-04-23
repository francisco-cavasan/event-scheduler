package Models

import "github.com/jinzhu/gorm"

type PetLocation struct {
	gorm.Model
	Pet      Pet
	Location Location
}
