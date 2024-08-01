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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	config.LoadConfig("../config/config_unittest.json")
	database.Connect(config.Config.MongoURI)
}

func TestCreatePost(t *testing.T) {
	// Set up the database connection
	postCollection := database.Client.Database(config.Config.Database).Collection("posts")
	postCollection.Drop(context.TODO()) // Clean up the collection before testing

	// Set up the Gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/posts", CreatePost)

	// Prepare the request payload
	post := models.Post{
		Title:   "Embracing Innovation in Product Management",
		Content: "Exploring the latest trends in product management and how innovation can drive success. Dive into new methodologies, tools, and strategies that are shaping the future of our field.",
		Author:  "Jane Doe, Product Management Expert",
	}
	jsonValue, _ := json.Marshal(post)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer your-secret-token")

	// Perform the request
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, recorder.Code)

	var responsePost models.Post
	err := json.Unmarshal(recorder.Body.Bytes(), &responsePost)
	assert.NoError(t, err)
	assert.Equal(t, post.Title, responsePost.Title)
	assert.Equal(t, post.Content, responsePost.Content)
	assert.Equal(t, post.Author, responsePost.Author)
}

func TestGetPosts(t *testing.T) {
	// Set up the database connection
	postCollection := database.Client.Database(config.Config.Database).Collection("posts")
	postCollection.Drop(context.TODO()) // Clean up the collection before testing

	// Insert a test post
	testPost := models.Post{
		ID:      primitive.NewObjectID(),
		Title:   "Mastering the Art of Product Development",
		Content: "Unlock the secrets to successful product development with actionable insights and real-world examples. Transform your ideas into market-ready products with confidence.",
		Author:  "John Smith, Senior Product Manager",
	}
	postCollection.InsertOne(context.TODO(), testPost)

	// Set up the Gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/posts", GetPosts)

	// Perform the request
	req, _ := http.NewRequest("GET", "/posts", nil)
	req.Header.Set("Authorization", "Bearer your-secret-token")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, recorder.Code)

	var responsePosts []models.Post
	err := json.Unmarshal(recorder.Body.Bytes(), &responsePosts)
	assert.NoError(t, err)
	assert.Len(t, responsePosts, 1)
	assert.Equal(t, testPost.Title, responsePosts[0].Title)
	assert.Equal(t, testPost.Content, responsePosts[0].Content)
	assert.Equal(t, testPost.Author, responsePosts[0].Author)
}

func TestGetPost(t *testing.T) {
	// Set up the database connection
	postCollection := database.Client.Database(config.Config.Database).Collection("posts")
	postCollection.Drop(context.TODO()) // Clean up the collection before testing

	// Insert a test post
	testPost := models.Post{
		ID:      primitive.NewObjectID(),
		Title:   "The Power of User-Centric Design",
		Content: "Discover how user-centric design can revolutionize your product strategy. Learn to prioritize user needs and create products that resonate and drive engagement.",
		Author:  "Emily Johnson, UX Designer and Product Specialist",
	}
	postCollection.InsertOne(context.TODO(), testPost)

	// Set up the Gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/posts/:id", GetPost)

	// Perform the request
	req, _ := http.NewRequest("GET", "/posts/"+testPost.ID.Hex(), nil)
	req.Header.Set("Authorization", "Bearer your-secret-token")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, recorder.Code)

	var responsePost models.Post
	err := json.Unmarshal(recorder.Body.Bytes(), &responsePost)
	assert.NoError(t, err)
	assert.Equal(t, testPost.Title, responsePost.Title)
	assert.Equal(t, testPost.Content, responsePost.Content)
	assert.Equal(t, testPost.Author, responsePost.Author)
}

func TestUpdatePost(t *testing.T) {
	// Set up the database connection
	postCollection := database.Client.Database(config.Config.Database).Collection("posts")
	postCollection.Drop(context.TODO()) // Clean up the collection before testing

	// Insert a test post
	testPost := models.Post{
		ID:      primitive.NewObjectID(),
		Title:   "Leveraging Data for Product Success",
		Content: "Harness the power of data to enhance your product development process. From analytics to user feedback, find out how data can inform and guide your decisions.",
		Author:  "Michael Brown, Data-Driven Product Manager",
	}
	postCollection.InsertOne(context.TODO(), testPost)

	// Set up the Gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/posts/:id", UpdatePost)

	// Prepare the request payload
	updatedPost := models.Post{
		Title:   "Future-Proofing Your Product Strategy",
		Content: "Stay ahead of the curve with future-proof strategies in product management. Learn to anticipate market trends and adapt your approach for sustained growth and success.",
		Author:  "Sarah Lee, Product Strategy Consultant",
	}
	jsonValue, _ := json.Marshal(updatedPost)
	req, _ := http.NewRequest("PUT", "/posts/"+testPost.ID.Hex(), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer your-secret-token")

	// Perform the request
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Verify the update in the database
	var responsePost models.Post
	err := postCollection.FindOne(context.TODO(), bson.M{"_id": testPost.ID}).Decode(&responsePost)
	assert.NoError(t, err)
	assert.Equal(t, updatedPost.Title, responsePost.Title)
	assert.Equal(t, updatedPost.Content, responsePost.Content)
}

func TestDeletePost(t *testing.T) {
	// Set up the database connection
	postCollection := database.Client.Database(config.Config.Database).Collection("posts")
	postCollection.Drop(context.TODO()) // Clean up the collection before testing

	// Insert a test post
	testPost := models.Post{
		ID:      primitive.NewObjectID(),
		Title:   "Agile Transformation in Product Management",
		Content: "Transform your product management approach with Agile methodologies. Learn how to foster collaboration, increase efficiency, and deliver high-quality products faster.",
		Author:  "David Green, Agile Coach and Product Leader",
	}
	postCollection.InsertOne(context.TODO(), testPost)

	// Set up the Gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE("/posts/:id", DeletePost)

	// Perform the request
	req, _ := http.NewRequest("DELETE", "/posts/"+testPost.ID.Hex(), nil)
	req.Header.Set("Authorization", "Bearer your-secret-token")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Verify the deletion in the database
	var responsePost models.Post
	err := postCollection.FindOne(context.TODO(), bson.M{"_id": testPost.ID}).Decode(&responsePost)
	assert.Error(t, err)
	assert.Equal(t, mongo.ErrNoDocuments, err)
}
