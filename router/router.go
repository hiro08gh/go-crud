package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-psql-api/controller"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run(":3001")
}

func router() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, "Hello World")
		})
		api.GET("/posts", controller.GetPosts)
		api.POST("/posts", controller.CreatePosts)
	}

	return r
}
