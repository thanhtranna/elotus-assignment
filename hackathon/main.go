package main

import (
	"github.com/gin-gonic/gin"

	"github.com/thanhtranna/elotus-assignment/hackathon/controllers"
	"github.com/thanhtranna/elotus-assignment/hackathon/database"
	"github.com/thanhtranna/elotus-assignment/hackathon/middlewares"
)

const (
	LimitMultipartMemory = 8
)

func main() {
	// Initialize Database
	database.Connect("root:password@tcp(mysql:3306)/hackathon_db?parseTime=true")
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = LimitMultipartMemory << 20 // 8 MiB

	api := router.Group("/api")
	{
		api.GET("/ping", controllers.Ping)
		api.POST("/user/login", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.POST("/upload", controllers.UploadFile)
		}
	}

	router.Static("/public", controllers.PathSaveImage)
	return router
}
