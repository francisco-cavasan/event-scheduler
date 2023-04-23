package Models

import "github.com/jinzhu/gorm"

type Location struct {
	gorm.Model
	State     string
	City      string
	Street    string
	Number    string
	Reference string
	Latitude  *float64 // may be null
	Longitude *float64 // may be null
}
