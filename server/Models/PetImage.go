package Models

import "github.com/jinzhu/gorm"

type PetImage struct {
	gorm.Model
	Pet   Pet
	Image Image
}
