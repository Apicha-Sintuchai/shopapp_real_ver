package handler

import (
	model "apicha/foodshop/Model"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Authhandler struct {
	db *gorm.DB
}

func NewAuthhandler(db *gorm.DB) *Authhandler {
	return &Authhandler{db: db}
}

func (h *Authhandler) Register(c *gin.Context) {
	var usernameandpassword model.Auth

	if err := c.BindJSON(&usernameandpassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	salt, err := bcrypt.GenerateFromPassword([]byte(usernameandpassword.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	usernameandpassword.Password = string(salt)

	if err := h.db.Create(&usernameandpassword).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "succesful",
	})
}

func (h *Authhandler) Login(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	secretket := os.Getenv("SECRET_KEY")
	var loginDetails model.Auth
	hmacSampleSecret := []byte(secretket)
	if err := c.BindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user model.Auth
	if err := h.db.Where("username = ?", loginDetails.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDetails.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
	})
}

func (h *Authhandler) Middleware() gin.HandlerFunc {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	secretket := os.Getenv("SECRET_KEY")
	return func(c *gin.Context) {
		hmacSampleSecret := []byte(secretket)
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			c.Abort()
			return
		}

		tokenString := authHeader[len("Bearer "):]
		tkn, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return hmacSampleSecret, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		if !tkn.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
