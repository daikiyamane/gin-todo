package todo

import (
	"gin-todo/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Controller is todo controller
type Controller struct{}

//Index action: Get /
func (ct Controller) Index(c *gin.Context) {
	var t models.Service
	todos, err := t.GetAll()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	}
}

//Create action: POST /users
func (ct Controller) Create(c *gin.Context) {
	var t models.Service
	_, err := t.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

//Show action: Get users/:id
func (ct Controller) Show(c *gin.Context) {
	id := c.Param("id")
	var t models.Service
	todo, err := t.GetOne(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.HTML(http.StatusOK, "edit.html", gin.H{"todo": todo})
	}
}

//Update action: PUT /update/:id
func (ct Controller) Update(c *gin.Context) {
	id := c.Param("id")
	var t models.Service
	_, err := t.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

// Delete action: DELETE /users/:id
func (ct Controller) Delete(c *gin.Context) {
	id := c.Param("id")
	var t models.Service

	if err := t.DeletByID(id); err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		log.Println(err)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
