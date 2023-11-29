package routes

import (
	controllers "github.com/Slightly-Techie/okr-api/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	// incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users/:id", controllers.GetUser())
}
