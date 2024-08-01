package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/VisarutJDev/social-media-api/database"
	"github.com/VisarutJDev/social-media-api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte("your_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Register godoc
//
//	@Summary		create user
//	@Description	create new user with username password
//	@Tags			user
//	@ID				Register
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.User		true	"register"
//	@Success		200		{object}	models.Response	"OK"
//	@Failure		400		{object}	models.Response	"Bad Request"
//	@Failure		401		{object}	models.Response	"Unauthorized"
//	@Failure		500		{object}	models.Response	"Internal Server Error"
//	@Router			/register [post]
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error(),
		})
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.Client.Database("social_media").Collection("users").FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&user)
	if err == nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			Error: "Username already exist",
		})
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "Username already exist"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: "Error while hashing password",
		})
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
		return
	}
	user.Password = string(hashedPassword)
	user.ID = primitive.NewObjectID()

	_, err = database.Client.Database("social_media").Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: "Error while creating user",
		})
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Message: "User registered successfully",
	})
	// c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login
//	@ID				Login
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			loginInput	body		models.LoginInput	true	"login"
//	@Success		200			{object}	models.AuthResponse	"OK"
//	@Failure		400			{object}	models.Response		"Bad Request"
//	@Failure		401			{object}	models.Response		"Unauthorized"
//	@Failure		500			{object}	models.Response		"Internal Server Error"
//	@Router			/login [post]
func Login(c *gin.Context) {
	var loginInput models.LoginInput
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error(),
		})
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := database.Client.Database("social_media").Collection("users").FindOne(context.Background(), bson.M{"username": loginInput.Username}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			Error: "Invalid username or password",
		})
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			Error: "Invalid username or password",
		})
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: loginInput.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: "Error while generating token",
		})
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while generating token"})
		return
	}

	c.JSON(http.StatusOK, models.AuthResponse{
		Token: tokenString,
	})
	// c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
