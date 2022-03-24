package main

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
}
