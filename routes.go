package main

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	router.Use(setUserStatus())
	router.GET("/", showIndexPage)
	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)
		userRoutes.GET("/logout", ensureLoggedIn(), logout)
	}

	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", getArticle)
		articleRoutes.GET("/create", ensureLoggedIn(), showArticleCreationPage)
		articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
	}
}
