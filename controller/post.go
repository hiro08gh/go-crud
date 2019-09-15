package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-crud/model"
)

func GetPosts(c *gin.Context) {
	db := model.GetDBConn()

	var posts []model.Post

	if err := db.Find(&posts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, &posts)
	}
}

//POST /api/projects
func CreatePosts(c *gin.Context) {
	db := model.GetDBConn()

	var posts model.Post

	posts.Title = string(c.Request.FormValue("Title"))

	c.Bind(&posts)
	db.Create(&posts)
	c.JSON(200, "ok")
}
