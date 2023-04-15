package server

import "github.com/jinzhu/gorm"

type Event struct {
	gorm.Model
	Name        string
	Description string
	Date        string
}
