package routes

import (
	"github.com/Slightly-Techie/okr-api/controllers"
	"github.com/gin-gonic/gin"
)

func ObjectiveRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/objective/create", controllers.CreateObjective())
	incomingRoutes.GET("/objective/get-all/:id", controllers.GetObjectives())
	incomingRoutes.PUT("/objective/update/:id", controllers.UpdateObjective())
	incomingRoutes.DELETE("/objective/delete/:id", controllers.DeleteObjective())
}
