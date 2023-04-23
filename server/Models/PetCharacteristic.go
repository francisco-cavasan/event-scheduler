package Models

import "github.com/jinzhu/gorm"

type PetCharacteristic struct {
	gorm.Model
	Pet            Pet
	Characteristic Characteristic
}
