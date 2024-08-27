package main

import (
	"fmt"
	"log"

	docs "github.com/HousewareHQ/backend-engineering-octernship/api"
	config "github.com/HousewareHQ/backend-engineering-octernship/cmd/configs"
	"github.com/HousewareHQ/backend-engineering-octernship/cmd/controllers/controller_v1"
	"github.com/HousewareHQ/backend-engineering-octernship/cmd/handlers"
	"github.com/HousewareHQ/backend-engineering-octernship/cmd/middleware"
	"github.com/HousewareHQ/backend-engineering-octernship/cmd/routes/routes_v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Authorization+Authentication service in Golang
// @verion 1.0.0
// description A backend API service in Golang that handles authorization and authentication for a web app

// @contact.name Shikhar Yadav
// @contact.email connectoshikhar10@gmail.com

// @host 127.0.0.1:10220
// @BasePath /api/v1

// @securityDefinitions.apikey Bearer
// @in Header
// @name Authorization

// @securityDefinitions.apikey Cookie
// @in Header
// @name Cookie

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Loading environment variables
	env := config.LoadENV()

	// setting gin_mode
	if env.GIN_MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// connecting to mongodb database
	mongodb, err := config.ConnectMongoDB(env)
	if err != nil {
		log.Fatal(err)
	}
	redisClient, err := config.ConnectRedis(env)
	if err != nil {
		log.Fatal(err)
	}

	// initializing database handlers which consists higher level functions for different operations.
	database := handlers.CreateDatabaseHandler(mongodb)
	cache := handlers.CreateCacheHandler(redisClient)

	// initializing jwt middleware
	jwt := middleware.CreateJWT(env, database, cache)

	// initializing user controllers
	userController := &controller_v1.UserController{
		Database: database,
		Cache:    cache,
		Jwt:      jwt,
	}

	// initializing organisation controllers
	adminController := &controller_v1.AdminController{
		Database: database,
	}

	// initializing gin router engine
	router := gin.New()
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// creating router group that has bash "/api/v1"
	v1Router := router.Group("/api/v1")

	// initializing user routes
	routes_v1.User(v1Router, userController, jwt)

	// initializing organisation routes
	routes_v1.Admin(v1Router, adminController, jwt)

	fmt.Println("Starting API WebServer At [:" + env.API_WEBSERVER_PORT + "]")
	// starting api webserver

	router.Run(":" + env.API_WEBSERVER_PORT)
}
