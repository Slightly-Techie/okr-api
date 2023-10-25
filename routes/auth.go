package routes

import(
	controllers "github.com/Slightly-Techie/okr-api/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/auth/login", controllers.Login())
}