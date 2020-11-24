package models

import (
	"gin-todo/db"
	"gin-todo/entity"
	"log"

	"github.com/gin-gonic/gin"
)

//Service procides user's behavior
type Service struct{}

//Todo is entity.Todo
type Todo entity.Todo

//GetAll gets all Todos
func (s Service) GetAll() ([]Todo, error) {
	db := db.GetDB()
	var t []Todo

	if err := db.Find(&t).Error; err != nil {
		return nil, err
	}

	return t, nil
}

//GetOne gets a Todo
func (s Service) GetOne(id string) (Todo, error) {
	db := db.GetDB()
	var t Todo
	if err := db.Where("id = ?", id).First(&t).Error; err != nil {
		return t, err
	}

	return t, nil
}

//CreateModel creates a user
func (s Service) CreateModel(c *gin.Context) (Todo, error) {
	db := db.GetDB()
	var t Todo

	if err := c.Bind(&t); err != nil {
		log.Println(err)
		return t, err
	}

	if err := db.Create(&t).Error; err != nil {
		log.Println(err)
		return t, err
	}

	return t, nil
}

//UpdateByID is update a User
func (s Service) UpdateByID(id string, c *gin.Context) (Todo, error) {
	db := db.GetDB()
	var t Todo
	if err := db.Where("id = ?", id).First(&t).Error; err != nil {
		return t, err
	}

	if err := c.Bind(&t); err != nil {
		return t, err
	}
	db.Save(&t)
	return t, nil
}

//DeletByID is delete a User
func (s Service) DeletByID(id string) error {
	db := db.GetDB()
	var t Todo

	if err := db.Where("id = ?", id).Delete(&t).Error; err != nil {
		return err
	}
	return nil
}
