package entity

import (
	"gorm.io/gorm"
)

//Todo is Model
type Todo struct {
	gorm.Model
	CreatedBy string `form:"name" gorm:"not null"`
	Content   string `form:"content" gorm:"not null"`
	Status    string `form:"status"`
}
