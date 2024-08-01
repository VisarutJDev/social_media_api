package controllers

import (
	"context"
	"net/http"

	"github.com/VisarutJDev/social-media-api/database"
	"github.com/VisarutJDev/social-media-api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreatePost     godoc
//
//	@Summary		Create Post
//	@Description	Create post by id
//	@ID				CreatePost
//	@Tags			post
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			post	body		models.Post		true	"Post data to be Created"
//	@Success		200		{object}	models.Post		"OK"
//	@Failure		400		{object}	models.Response	"Bad Request"
//	@Failure		401		{object}	models.Response	"Unauthorized"
//	@Failure		500		{object}	models.Response	"Internal Server Error"
//	@Router			/posts [post]
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error(),
		})
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.ID = primitive.NewObjectID()
	_, err := database.Client.Database("social_media").Collection("posts").InsertOne(context.Background(), post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, post)
}

// GetPosts godoc
//
//	@Summary		Get Posts
//	@Description	Get posts
//	@ID				GetPosts
//	@Tags			post
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Post		"OK"
//	@Failure		400	{object}	models.Response	"Bad Request"
//	@Failure		401	{object}	models.Response	"Unauthorized"
//	@Failure		500	{object}	models.Response	"Internal Server Error"
//	@Router			/posts [get]
func GetPosts(c *gin.Context) {
	findOptions := options.Find()
	// Sort by `price` field descending
	findOptions.SetSort(bson.D{{Key: "_id", Value: -1}})

	cursor, err := database.Client.Database("social_media").Collection("posts").Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error(),
		})
		return
	}
	var posts []models.Post
	if err = cursor.All(context.Background(), &posts); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// GetPost godoc
//
//	@Summary		Get Post
//	@Description	Get post by id
//	@ID				GetPost
//	@Tags			post
//	@Security		Bearer
//	@iD				GetPost
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"id of post to be get"
//	@Success		200	{object}	models.Post		"OK"
//	@Failure		400	{object}	models.Response	"Bad Request"
//	@Failure		401	{object}	models.Response	"Unauthorized"
//	@Failure		500	{object}	models.Response	"Internal Server Error"
//	@Router			/posts/{id} [get]
func GetPost(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	var post models.Post
	err := database.Client.Database("social_media").Collection("posts").FindOne(context.Background(), bson.M{"_id": objID}).Decode(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: "Post not found",
		})
		// c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

// UpdatePost godoc
//
//	@Summary		Update Post
//	@Description	Update post by id
//	@ID				UpdatePost
//	@Tags			post
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"id of post to be updated"
//	@Param			post	body		models.Post		true	"Post data to be updated"
//	@Success		200		{object}	models.Response	"OK"
//	@Failure		400		{object}	models.Response	"Bad Request"
//	@Failure		401		{object}	models.Response	"Unauthorized"
//	@Failure		500		{object}	models.Response	"Internal Server Error"
//	@Router			/posts/{id} [put]
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error(),
		})
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := database.Client.Database("social_media").Collection("posts").UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": post})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Message: "Post updated successfully",
	})
	// c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

// DeletePost godoc
//
//	@Summary		Delete Post
//	@Description	Delete post by id
//	@Tags			post
//	@Security		Bearer
//	@ID				DeletePost
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"id of post to be deleted"
//	@Success		200	{object}	models.Response	"OK"
//	@Failure		400	{object}	models.Response	"Bad Request"
//	@Failure		401	{object}	models.Response	"Unauthorized"
//	@Failure		500	{object}	models.Response	"Internal Server Error"
//	@Router			/posts/{id} [delete]
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := database.Client.Database("social_media").Collection("posts").DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error(),
		})
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Message: "Post deleted successfully",
	})
	// c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
