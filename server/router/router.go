package router

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"server/internal/user"
	"server/internal/server"
)


func InitRouter(userHandler *user.Handler, serverHandler *server.Handler) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:5173"}, // Update to match the origin you want to allow
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
    }))

	router.POST("/createUser", userHandler.CreateUser)
	router.POST("/login", userHandler.Login)
	router.POST("/logout", userHandler.Logout)

	router.POST("/server/createServer", serverHandler.CreateServer)

	return router
}

func Start(router *gin.Engine, addr string) error {
	return router.Run(addr)
}