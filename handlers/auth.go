package handlers

import (
	"api_assignment/middleware"
	"api_assignment/models"
	"api_assignment/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"log"

)

var db *gorm.DB

func init() {
	var err error
	dsn := "root:root123@tcp(db:3306)/api_assignment?charset=utf8mb4&parseTime=True&loc=Local"
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			// If connection is successful, break out of the loop
			log.Println("Connected to the database")
			break
		}
		log.Printf("Failed to connect to database. Retrying... (attempt %d/10)", i+1)
		time.Sleep(1 * time.Second) 
	}

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})
}

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", Register)
	r.POST("/login", Login)
	r.POST("/revoke", RevokeTokenHandler)
	r.GET("/protected", middleware.TokenAuthMiddleware(), ProtectedEndpoint)
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	var dbUser models.User
	if err := db.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	accessToken, refreshToken, _ := services.GenerateTokens(user.Email)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func RevokeTokenHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing token"})
		return
	}
	services.RevokeToken(token)
	c.JSON(http.StatusOK, gin.H{"message": "Token revoked"})
}

func ProtectedEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You have accessed a protected endpoint!"})
}
