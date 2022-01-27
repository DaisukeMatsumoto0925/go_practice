package infrastructure

import (
	gin "github.com/gin-gonic/gin"

	"clean_arc/interfaces/controllers"
)

func NewRouter(controller *controllers.UserController) *gin.Engine {
	router := gin.Default()

	router.POST("/users", func(c *gin.Context) { controller.Create(c) })
	router.GET("/users", func(c *gin.Context) { controller.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { controller.Show(c) })

	return router
}
