package entity

import (
	"gorm.io/gorm"
)

//Todo is Model
type Todo struct {
	gorm.Model
	CreatedBy string `form:"name" gorm:"NOT NULL"`
	Content   string `form:"content" gorm:"NOT NULL"`
	Status    string `form:"status"`
}
