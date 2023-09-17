package main

import (
	"log"

	"github.com/Slightly-Techie/okr-api/database"
	"github.com/Slightly-Techie/okr-api/models"
	"github.com/Slightly-Techie/okr-api/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine) {
	r.GET("/", routes.DefaultHandler)
	r.POST("/test", routes.TestHandler)
}

func main() {
	db := database.GetConnection()
	db.AutoMigrate(models.Test{})

	router := gin.Default()
	setupRouter(router)
	err := router.Run(":5000")
	if err != nil {
		log.Fatalf("gin Run error: %s", err)
	}
}
