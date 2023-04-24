package Models

import "github.com/jinzhu/gorm"

type Pet struct {
	gorm.Model
	Name        string
	Description string
	Age         string
	LostDate    string
	FoundDate   string
	Owner       User
	Image_url   string
	// Images      []Image
}
