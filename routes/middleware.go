
package routes

import (
	"net/http"
	"log"
	"github.com/Slightly-Techie/okr-api/utils"
	"github.com/gin-gonic/gin"
)

func AuthenticationRequired(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		log.Fatal("TOKEN ERROR", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authentication required."})
		return
	}

	//var claims *jwt.MapClaims

	_, err = utils.DecodeJWT(token, JWT_SECRET)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid token"})
		return
	}
	c.Next()
}