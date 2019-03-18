package models

import (
	"github.com/jinzhu/gorm"

	"github.com/afdolriski/restful-test/lib/common"
)

// Image model struct
type Image struct {
	gorm.Model
	Name string
}

// Serialize as json
func (user *Image) Serialize() common.JSON {
	return common.JSON{
		"name": user.Name,
	}
}
