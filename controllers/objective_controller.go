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

func GetObjective() gin.HandlerFunc {
	return func(c *gin.Context) {

		var objectives models.Objective
		objId := c.Param("id")

		if err := database.GetItem(objectives, objId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": &objectives})
	}
}

func GetObjectives() gin.HandlerFunc {
	return func(c *gin.Context) {
		var objectives []models.Objective
		userId := c.Param("id")

		if err := database.GetItem(objectives, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": &objectives})
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

		if err := database.UpdateItem(&models.Objective{}, &objective, "objectiveId", objId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, "objective created successfully")
	}
}

func DeleteObjective() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model models.Objective
		objId := c.Param("id")

		if err := database.DeleteItem(model, "objectiveId", objId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "objective deleted successfully"})
	}
}
