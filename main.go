package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"main/browser"
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
	err := router.Run()
	if err != nil {
		return
	}

	// The browser will bw available in "/" endpoint and ready for use
	browser, err := browser.NewManager()
	if err != nil {
		panic(err)
	}
	defer browser.Close()

	router.GET("/", func(c *gin.Context) {
		page, err := browser.Navigate("https://www.google.com")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer page.Close()

		content, err := page.Content()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
	})

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}


}