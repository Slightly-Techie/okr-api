package routes

import (
	"net/http"
	"log"
	"time"
	"github.com/Slightly-Techie/okr-api/database"
	"github.com/Slightly-Techie/okr-api/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
	"github.com/Slightly-Techie/okr-api/utils"
)

const CLIENT_ID string = "YOUR_CLIENT_ID"
const JWT_SECRET string = "something else"

func DefaultHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func LoginHandler(c *gin.Context)  {
	var loginInfo models.LoginInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
      	log.Fatalf("Couldn't convert token to valid struct", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	payload, err := idtoken.Validate(c, loginInfo.Credential, CLIENT_ID)
	if err != nil {
		log.Fatalf("Could not validate sign in token", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JWT."})
		return
	}

	// create a JWT for the app and send it back to the client for future requests
	tokenString, err := utils.GenerateJWT(payload.Subject, JWT_SECRET)
	if err != nil {
		log.Fatalf("Failed to create JWT", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong completing your sign in."})
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour), // Cookie expiration time
		HttpOnly: true,                           // Cookie is not accessible via JavaScript
	}

	// Set the cookie in the response
	http.SetCookie(c.Writer, cookie)
	c.Status(http.StatusOK)
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


