package server

import "github.com/jinzhu/gorm"

type UserEvent struct {
	gorm.Model
	User  User
	Event Event
}
