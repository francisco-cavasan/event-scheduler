package Models

import "github.com/jinzhu/gorm"

type Characteristic struct {
	gorm.Model
	Name  string
	Value string
}
