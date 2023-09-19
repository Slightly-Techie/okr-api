package main

import (
	"log"
	"os"

	"github.com/Slightly-Techie/okr-api/database"
	"github.com/Slightly-Techie/okr-api/models"
	"github.com/Slightly-Techie/okr-api/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine) {
	r.GET("/", routes.DefaultHandler)
	r.POST("/test", routes.TestCreateHandler)
	r.GET("/get", routes.TestGetHandler)
}

func main() {
	database.InitDB()
	db := database.GetDB()
	db.AutoMigrate(models.Test{})

	router := gin.Default()
	setupRouter(router)
	err := router.Run(os.Getenv(":5000"))
	if err != nil {
		log.Fatalf("gin Run error: %s", err)
	}
}
