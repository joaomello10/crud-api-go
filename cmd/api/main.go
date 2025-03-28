package main

import (
	"fmt"
	"gorm/cmd/initializers"
	"gorm/cmd/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB
var products = []models.Note{}

func initApp() {
	//initializers.LoadEnvVars()
	DB = initializers.ConnectToDB()
}

//var DB *gorm.DB

func createPost(c *gin.Context) {
	var post models.Note

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
	fmt.Println("Post added successfully")
}

func getPosts(c *gin.Context) {
	var posts []models.Note
	DB.Find(&posts)
	c.JSON(200, posts)
}

func deletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Note
	DB.First(&post, id)

	if err := DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"post not found, Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Post deleted successfully"})
}

func updatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Note

	if err := DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"post not found, Error": err.Error()})
		return
	}

	c.BindJSON(&post)
	DB.Save(&post)

	c.JSON(200, gin.H{"message": "post updated successfully", "post": post})
}

func getPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Note

	if err := DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"post not found, Error": err.Error()})
		return
	}

	c.JSON(200, post)
}

func main() {
	initApp()

	log.Println("Connected to PostgreSQL successfully!")
	log.Println(products)

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	r.Use(cors.New(config))
	r.GET("/notes/:id", getPost)
	r.GET("/notes", getPosts)
	r.POST("/new", createPost)
	r.DELETE("/notes/:id", deletePost)
	r.PUT("/notes/:id", updatePost)
	r.Run(":8080")
}
