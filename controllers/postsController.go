package controllers

import (
	"crud-app/initializers"
	"crud-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Body, Body: body.Title}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	initializers.DB.First(&post, id)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	initializers.DB.First(&post, id)
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(http.StatusOK, gin.H{
		"post": "Post deleted",
	})
}
