package controllers

import(
	"log"
	"net/http"
	"time"
	"strings"
	"os"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
	"github.com/go-playground/validator/v10"
	"github.com/Slightly-Techie/okr-api/helpers"
	"github.com/Slightly-Techie/okr-api/models"
	"github.com/Slightly-Techie/okr-api/database"

)
var client_id string = os.Getenv("CLIENT_ID")

var validate = validator.New()

func Login() gin.HandlerFunc{
	log.Println("Here")
	return func(c *gin.Context){
		var loginInfo models.LoginInfo
		var user models.User
		var foundUser models.User

		db := database.GetDB()

		if err := c.ShouldBindJSON(&loginInfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		payload, err := idtoken.Validate(c, loginInfo.Credential, client_id)

		claims := payload.Claims
		log.Println("Payload: ", claims["email"])
		if err != nil {
			log.Fatalf("Could not validate sign in token: %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JWT."})
			return
		}

		email := claims["email"].(string)
		first_name := claims["given_name"].(string)
		last_name := claims["family_name"].(string)
		uid := claims["sub"].(string)

		err = db.First(&foundUser, "email = ?", claims["email"]).Error
		log.Println("Found User: ", foundUser)
		
		if err != nil && !strings.Contains(err.Error(), "not found"){
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, foundUser)
		token, refreshToken := helpers.GenerateAllTokens(email, first_name, last_name, uid)

		if err != nil && strings.Contains(err.Error(), "not found") {
			user = models.User{
				Email: email,
				FirstName: first_name,
				LastName: last_name,
				Uid: uid,
			}
			user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			
			user.Token = token
			user.RefreshToken = refreshToken
			db.Create(&user)

			c.JSON(http.StatusCreated, &user)
			return
		}

		foundUser.Token = token
		foundUser.RefreshToken = refreshToken
		foundUser.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		db.Save(foundUser)
		c.JSON(http.StatusOK, foundUser)
		return
	}
}

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		db := database.GetDB()
		userId := c.Param("id")

		var user models.User
		
		err := db.First(&user, userId).Error
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

