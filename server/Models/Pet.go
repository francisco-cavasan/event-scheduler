package Models

import "github.com/jinzhu/gorm"

type Pet struct {
	gorm.Model
	Name            string
	Description     string
	Age             string
	Characteristics []Characteristic
	LostDate        string
	Owner           User
	Images          []Image
}
