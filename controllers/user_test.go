package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VisarutJDev/social-media-api/config"
	"github.com/VisarutJDev/social-media-api/database"
	"github.com/VisarutJDev/social-media-api/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	config.LoadConfig("../config/config_unittest.json")
	database.Connect(config.Config.MongoURI)
}

func TestRegister(t *testing.T) {
	userCollection := database.Client.Database(config.Config.Database).Collection("users")
	userCollection.Drop(context.TODO())

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/register", Register)

	user := models.User{
		Username: "testuser",
		Password: "password123",
	}
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	var response map[string]string
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "User registered successfully", response["message"])
}

func TestLogin(t *testing.T) {
	userCollection := database.Client.Database(config.Config.Database).Collection("users")
	userCollection.Drop(context.TODO())

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	testUser := models.User{
		Username: "testuser",
		Password: string(hashedPassword),
	}
	userCollection.InsertOne(context.TODO(), testUser)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/login", Login)

	loginInput := models.LoginInput{
		Username: "testuser",
		Password: "password123",
	}
	jsonValue, _ := json.Marshal(loginInput)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var response map[string]string
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response["token"])
}
