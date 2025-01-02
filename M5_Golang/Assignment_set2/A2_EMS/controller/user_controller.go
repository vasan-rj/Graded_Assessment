package controller

import (
	"ecommerce/model"
	"ecommerce/service"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *service.AccountService
}

// NewUserController initializes and returns a UserController instance.
func NewUserController(userService *service.AccountService) *UserController {
	return &UserController{UserService: userService}
}

// Register handles user registration.
func (controller *UserController) Register(c *gin.Context) {
	var user model.Account
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.UserService.RegisterAccount(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login handles user authentication and token generation.
func (controller *UserController) Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authenticate user
	user, err := controller.UserService.ValidateCredentials(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT token
	token, err := generateJWT(user.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating JWT token"})
		return
	}

	// Responding to client with token
	c.JSON(http.StatusOK, gin.H{
		"message": "Authentication successful",
		"token":   token,
	})
}

// generateJWT generates a JWT token for the given username.
func generateJWT(username string) (string, error) {
	// Define the token expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, which includes the username and expiry time
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    "ecommerce-inventory",
		Subject:   username,
	}

	// Create the token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("secretkey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
