package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "main/docs"
)

// @title Blog API
// @version 1.0
// @description API for managing blog posts.
// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/blog", getting)
	router.POST("/blog", posting)

	err := router.Run()
	if err != nil {
		return
	}
}

// getting godoc
// @Summary Get blog
// @Description Get blog information
// @Tags blog
// @Produce json
// @Success 200 {object} map[string]string
// @Router /blog [get]
func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Getting the blog through the GET method",
	})
}

// posting godoc
// @Summary Create blog
// @Description Create a blog post
// @Tags blog
// @Produce json
// @Success 200 {object} map[string]string
// @Router /blog [post]
func posting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Posting the blog through POST method",
	})
}