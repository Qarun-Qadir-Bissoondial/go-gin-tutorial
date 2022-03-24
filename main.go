package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")
	router.GET("/", showIndexPage)

	err := router.Run()
	if err != nil {
		log.Fatalln("Could not start application!")
	}
}
