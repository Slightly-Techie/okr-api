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
	// r.POST("/test", routes.TestCreateHandler)
	// r.GET("/get", routes.TestGetItemsHandler)
	// r.GET("/getItem", routes.TestGetItemHandler)
	// r.PUT("/update", routes.TestUpdateItemHandler)
	// r.DELETE("/delete", routes.TestDeleteItemHandler)
}

func main() {
	database.InitDB()
	db := database.GetDB()
	db.AutoMigrate(models.Test{}, models.Company{}, models.Objective{}, models.KeyResult{})

	router := gin.Default()
	setupRouter(router)
	err := router.Run(":5000")
	if err != nil {
		log.Fatalf("gin Run error: %s", err)
	}
}
