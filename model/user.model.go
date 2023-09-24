// models/user.go
package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name        string
	PhoneNumber string `gorm:"column:phoneNumber"`
	Password    string
}
