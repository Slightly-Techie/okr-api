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

// func TestCreateHandler(c *gin.Context) {
// 	var data models.Test
// 	err := c.ShouldBind(&data)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"success": false})
// 		return
// 	}
// 	database.CreateItem(&data)
// 	c.JSON(http.StatusOK, gin.H{"success": data})
// }

// func TestGetItemsHandler(c *gin.Context) {
// 	var data []models.Test
// 	_, err := database.GetAllItems(&data, "name", "dexter")

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
// }

// func TestGetItemHandler(c *gin.Context) {
// 	var data models.Test
// 	_, err := database.GetItem(&data, "id", "1")

// 	if err != nil {
// 		c.JSON(http.StatusBadGateway, gin.H{"success": "false"})
// 		return
// 	}
// 	if data.ID == 0 {
// 		c.JSON(http.StatusOK, gin.H{"success": true, "data": ""})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
// }

// func TestUpdateItemHandler(c *gin.Context) {
// 	var data models.Test
// 	err := database.UpdateItem(data, "id", "4", "name", "diego")

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"success": false})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"success": true, "message": "updated successfully"})
// }

// func TestDeleteItemHandler(c *gin.Context) {
// 	var data models.Test

// 	err := database.DeleteItem(data, "id", "1")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"success": true, "message": "deleted successfully"})

// }
