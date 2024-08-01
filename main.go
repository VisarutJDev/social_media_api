package main

import (
	"github.com/VisarutJDev/social-media-api/config"
	"github.com/VisarutJDev/social-media-api/database"
	"github.com/VisarutJDev/social-media-api/routes"

	"github.com/gin-gonic/gin"

	"github.com/VisarutJDev/social-media-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title Social media API
//	@version 1.0
//	@termsOfService	https://portfolio-6b550.web.app/

//	@contact.name API Support
//	@contact.url https://portfolio-6b550.web.app/
//	@contact.email visarutjwork@gmail.com

//	@license.name Apache 2.0
//	@license.url http://www.apache.org/licenses/LICENSE-2.0.html

//	@schemes http
//	@BasePath /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Social media API"
	docs.SwaggerInfo.Description = "This is a sample server Social media server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	config.LoadConfig("config/config_local.json")
	database.Connect(config.Config.MongoURI)

	router := gin.Default()
	// router.Use(middlewares.TokenAuthMiddleware())
	routes.InitRoutes(router)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
