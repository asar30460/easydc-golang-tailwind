package router

import (
	"server/internal/server"
	"server/internal/user"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func InitRouter(userHandler *user.Handler, serverHandler *server.Handler) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Update to match the origin you want to allow
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	}))

	router.POST("/createUser", userHandler.CreateUser)
	router.POST("/login", userHandler.Login)
	router.POST("/logout", userHandler.Logout)

	router.GET("/server/handleWs", serverHandler.HandleWS)

	router.POST("/server/createServer", serverHandler.CreateServer)
	router.GET("/server/getServers", serverHandler.GetServer)
	router.POST("/server/joinServer", serverHandler.JoinServer)

	router.POST("/server/:server_id/createChannel", serverHandler.CreateChannel)
	router.GET("/server/:server_id/getChannels", serverHandler.GetChannel)
	router.GET("/server/:server_id/getMembers", serverHandler.GetMember)

	// This is for get historys msg in the given channel.
	// Since channel_id is unique, so we don't specify url.
	router.POST("/server/createMsg", serverHandler.CreateMsg)
	router.POST("/server/getHistoryMsgs", serverHandler.GetHistorysMsg)

	return router
}

func Start(router *gin.Engine, addr string) error {
	return router.Run(addr)
}
