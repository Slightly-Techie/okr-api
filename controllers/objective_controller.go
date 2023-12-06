package controllers

import (
	"net/http"

	"github.com/Slightly-Techie/okr-api/database"
	"github.com/Slightly-Techie/okr-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateObjective() gin.HandlerFunc {

	return func(c *gin.Context) {

		var objective models.Objective

		if err := c.ShouldBindJSON(&objective); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := uuid.New()
		data := models.Objective{
			ObjectiveId: id.String(),
			Title:       objective.Title,
			Description: objective.Description,
			UserId:      objective.UserId,
			Assignee:    objective.Assignee,
			CompanyId:   objective.CompanyId,
		}

		if err := database.CreateItem(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "objective created successfully")

	}
}

func GetObjectives() gin.HandlerFunc {
	return func(c *gin.Context) {
		var objectives *[]models.Objective
		userId := c.Param("id")

		data, err := database.GetItems(objectives, "user_id", userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func UpdateObjective() gin.HandlerFunc {
	return func(c *gin.Context) {
		var objective models.Objective
		objId := c.Param("id")

		if err := c.ShouldBindJSON(&objective); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		update := map[string]interface{}{"title": &objective.Title, "description": &objective.Description, "assignee": &objective.Assignee}

		if err := database.UpdateItem(&models.Objective{}, update, "objective_id", objId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "objective updated successfully")
	}
}

func DeleteObjective() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model models.Objective
		objId := c.Param("id")

		if err := database.DeleteItem(model, "objective_id", objId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "objective deleted successfully"})
	}
}
