package Controllers

import (
	"net/http"
	"where_my_pet_at/server/Models"
	"where_my_pet_at/server/Services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	DB *gorm.DB
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginPayload LoginPayload
	err := c.BindJSON(&loginPayload)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user := Models.User{}

	err = ac.DB.Debug().Model(Models.User{}).Where("email = ?", loginPayload.Email).Take(&user).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	token, err := ac.SignIn(user.Email, loginPayload.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ac *AuthController) SignIn(email string, password string) (string, error) {
	var err error

	user := Models.User{}

	err = ac.DB.Debug().Model(Models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}

	err = Models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := Services.CreateToken(uint32(user.ID))

	return token, err
}
