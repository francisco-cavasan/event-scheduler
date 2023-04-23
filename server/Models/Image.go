package Models

import "github.com/jinzhu/gorm"

type Image struct {
	gorm.Model
	Name string
	Url  string
}
