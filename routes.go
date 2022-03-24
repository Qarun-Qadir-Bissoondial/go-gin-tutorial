package main

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	router.GET("/", showIndexPage)
	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/register", showRegistrationPage)
		userRoutes.POST("/register", register)
	}

	router.GET("/article/view/:article_id", getArticle)
}
