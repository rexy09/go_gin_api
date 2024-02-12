package controllers

import (
	// "fmt"
	"go_gin_api/database"
	"go_gin_api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)
	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := database.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return

	}

	c.JSON(200, gin.H{"post": post})

}

func PostList(c *gin.Context) {
	// id := c.Param("id")
	// id := c.Query("id")
	id := c.DefaultQuery("id", "")
	// fmt.Println("Hello Params")
	// fmt.Println(id)
	// fmt.Println("Bye Params")
	var posts []models.Post
	var result *gorm.DB

	// if id != "" && title != "" {
	//     // If both id and title are provided, fetch posts with both conditions
	//     result = database.DB.Where("id = ? AND title = ?", id, title).Find(&posts)
	// } else
	// else if title != "" {
	//     // If only title is provided, fetch posts with the specified title
	//     result = database.DB.Where("title = ?", title).Find(&posts)
	// }
	if id != "" {
		// If only id is provided, fetch posts with the specified id
		result := database.DB.Where(map[string]interface{}{"id": id}).Find(&posts)
		if result.Error != nil {
			c.Status(400)
			return

		}

	} else {
		// If neither id nor title is provided, fetch all posts
		result = database.DB.Find(&posts)
		if result.Error != nil {
			c.Status(400)
			return
	
		}
	}


	c.JSON(200, gin.H{"data": posts})

}
