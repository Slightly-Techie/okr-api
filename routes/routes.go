package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}
