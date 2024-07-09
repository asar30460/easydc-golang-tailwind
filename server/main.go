package main

import (
	"fmt"
	"server/db"
	"server/internal/user"
	"server/internal/server"
	"server/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Initilize db in main. ", err)
	}
	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	// For server HTTP and WS
	serverRep := server.NewRepository(dbConn.GetDB())
	hub:= server.NewHub()
	serverHandler := server.NewHandler(serverRep, hub)

	// For server WS

	r := router.InitRouter(userHandler, serverHandler)
	router.Start(r, ":8080")
}
