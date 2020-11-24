package server

import (
	"github.com/gin-gonic/gin"

	todo "gin-todo/controller"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	u := r.Group("/")
	{
		ctrl := todo.Controller{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.POST("/update/:id", ctrl.Update)
		u.POST("/delete/:id", ctrl.Delete)
	}

	return r
}
