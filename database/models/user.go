package models

import (
	"github.com/jinzhu/gorm"

	"github.com/afdolriski/restful-test/lib/common"
)

// User model struct
type User struct {
	gorm.Model
	Name        string
	Email       string
	PhoneNumber string
}

// Serialize as json
func (user *User) Serialize() common.JSON {
	return common.JSON{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"phone": user.PhoneNumber,
	}
}

func (user *User) Read(m common.JSON) {
	user.ID = uint(m["id"].(float64))
	user.Name = m["name"].(string)
	user.Email = m["email"].(string)
	user.PhoneNumber = m["phone"].(string)
}
