package routes

import (
	"net/http"

	"github.com/Slightly-Techie/okr-api/database"
	"github.com/Slightly-Techie/okr-api/models"
	"github.com/gin-gonic/gin"
)

func DefaultHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func TestCreateHandler(c *gin.Context) {
	var data models.Test
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	database.CreateItem(&data)
	c.JSON(http.StatusOK, gin.H{"success": data})
}

func TestGetHandler(c *gin.Context) {
	var data []models.Test
	_, err := database.GetItems(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}
