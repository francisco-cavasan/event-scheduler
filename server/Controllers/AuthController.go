package Controllers

import (
	"net/http"

	"where_my_pet_at/server/Models"
	"where_my_pet_at/server/utils/Token"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

var sampleSecretKey = []byte("SecretYouShouldHide")

type Credentials struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (ac *AuthController) Register(c *gin.Context) {
	var creds Credentials

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := Models.User{}

	u.Email = creds.Email
	u.Password = creds.Password

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	var input Credentials

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := Models.User{}

	u.Email = input.Email
	u.Password = input.Password

	token, err := Models.LoginCheck(u.Email, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func CurrentUser(c *gin.Context) {

	user_id, err := Token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := Models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func (ac *AuthController) Logout(c *gin.Context) {
}
