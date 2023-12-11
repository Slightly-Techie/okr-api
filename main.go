package main

import (
	"log"
	"net/http"

	"github.com/Slightly-Techie/okr-api/database"
	"github.com/Slightly-Techie/okr-api/models"
	"github.com/Slightly-Techie/okr-api/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine) {
	// r.Use(routes.AuthenticationRequired())
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true})
	})
}

func main() {
	database.InitDB()
	db := database.GetDB()
	db.AutoMigrate(models.User{}, models.Objective{})

	router := gin.Default()
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.ObjectiveRoutes(router)
	setupRouter(router)
	err := router.Run(":5000")
	if err != nil {
		log.Fatalf("gin Run error: %s", err)
	}
}
