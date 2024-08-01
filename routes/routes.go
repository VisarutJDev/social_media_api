package routes

import (
	"github.com/VisarutJDev/social-media-api/controllers"
	"github.com/VisarutJDev/social-media-api/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/healthcheck", controllers.HealthCheckHandler)
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middlewares.AuthMiddleware())
	{
		protectedRoutes.POST("/posts", controllers.CreatePost)
		protectedRoutes.GET("/posts", controllers.GetPosts)
		protectedRoutes.GET("/posts/:id", controllers.GetPost)
		protectedRoutes.PUT("/posts/:id", controllers.UpdatePost)
		protectedRoutes.DELETE("/posts/:id", controllers.DeletePost)
	}
}
