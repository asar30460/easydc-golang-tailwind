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
	// For user HTTP
	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	// For server HTTP and WS
	serverRep := server.NewRepository(dbConn.GetDB())
	serverSvc := server.NewService(serverRep)
	serverHandler := server.NewHandler(serverSvc)

	r := router.InitRouter(userHandler, serverHandler)
	router.Start(r, ":8080")
}
