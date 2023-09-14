package main

import (
	"log"

	"github.com/Slightly-Techie/okr-api/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine) {
	r.GET("/", routes.DefaultHandler)
}

func main() {
	router := gin.Default()
	setupRouter(router)
	err := router.Run(":5000")
	if err != nil {
		log.Fatalf("gin Run error: %s", err)
	}
}
